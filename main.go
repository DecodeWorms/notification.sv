package main

import (
	"github.com/DecodeWorms/notification.sv/config"
	serverutils "github.com/DecodeWorms/notification.sv/server-utils"
)

var c config.Config

func main() {
	c = config.ImportConfig(config.OSSource{})
	_ = serverutils.SetUpSmtpServer(c.Host, c.Port)
	router := serverutils.SetUpRouter()

	serverutils.StartServer(router)
}
