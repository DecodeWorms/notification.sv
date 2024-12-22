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
	smtp, err := serverutils.SetUpSmtpServer(c)
	if err != nil {
		log.Printf("error setting up smtp server %v", err)
	}
	sub, err := pulsar.NewPulsarClient(c.PulsarUrl)
	if err != nil {
		log.Printf("error connecting to Pulsar client %v", err)
	}

	//Initialize event listener
	subscriber := serverutils.SetUpSubscriber(sub, smtp)

	//Call event listener for user
	subscriber.SubscribeToVerifyEmail()
	subscriber.SubscribeToWelcomeEmail()
	subscriber.SubscribeToSendForgotPasswordCode()
	subscriber.SubscribeToSuccessfulResetPassword()
	subscriber.SubscribeToSuccessfulResetChangePassword()
	subscriber.SubscribeToSuccessfulAidCreation()
	subscriber.SubscribeToSuccessfulAidUpdating()
	subscriber.SubscribeToSuccessfulAidDeleted()

	//Call event listener for company
	subscriber.SubscribeToCompanyVerifyEmail()

	router := serverutils.SetUpRouter()

	serverutils.StartServer(router, subscriber)
}
