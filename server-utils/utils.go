package serverutils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/DecodeWorms/notification.sv/config"
	"github.com/DecodeWorms/notification.sv/handler"
	"github.com/DecodeWorms/notification.sv/notify"
	"github.com/DecodeWorms/pulsify/pulsar"
	"github.com/gin-gonic/gin"
)

func SetUpSmtpServer(c config.Config) (*notify.SmtpServer, error) {
	return notify.NewSmtServer(c)
}

func SetUpSubscriber(sub *pulsar.PulsarClient, smtp *notify.SmtpServer) handler.Subscriber {
	return handler.NewSubscriber(sub, *smtp)
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func StartServer(router *gin.Engine, sub handler.Subscriber) {
	//var c config.Config
	var c = config.ImportConfig(config.OSSource{})
	interruptHandler := make(chan os.Signal, 1)
	signal.Notify(interruptHandler, syscall.SIGTERM, syscall.SIGINT)

	addr := fmt.Sprintf(":%s", c.ServicePort)
	go func(addr string) {
		log.Println(fmt.Sprintf("Notification.sv API service running on %v. Environment=%s", addr, c.AppEnv))
		if err := http.ListenAndServe(addr, router); err != nil {
			log.Printf("Error starting server: %v", err)
		}
	}(addr)

	<-interruptHandler
	log.Println("Closing application...")

	//Clean up resources
	sub.Shutdown()

}
