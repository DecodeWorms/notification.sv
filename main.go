package main

import (
	"log"

	"github.com/DecodeWorms/notification.sv/config"
	serverutils "github.com/DecodeWorms/notification.sv/server-utils"
	"github.com/DecodeWorms/pulsify/pulsar"
)

var c config.Config

func main() {
	c = config.ImportConfig(config.OSSource{})
	smtp := serverutils.SetUpSmtpServer(c)
	sub, err := pulsar.NewPulsarClient(c.PulsarUrl)
	if err != nil {
		log.Printf("error connecting to Pulsar client %v", err)
	}

	//Initialize event listener
	subscriber := serverutils.SetUpSubscriber(sub, smtp)

	//Call event listener
	subscriber.SubscribeToVerifyEmail()

	router := serverutils.SetUpRouter()

	serverutils.StartServer(router, subscriber)
}
