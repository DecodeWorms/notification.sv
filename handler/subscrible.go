package handler

import (
	"encoding/json"
	"log"

	"github.com/DecodeWorms/notification.sv/constant"
	"github.com/DecodeWorms/notification.sv/models"
	"github.com/DecodeWorms/notification.sv/notify"

	"github.com/DecodeWorms/pulsify/pulsar"
)

type Subscriber struct {
	sub  *pulsar.PulsarClient
	smtp *notify.SmtpServer
}

func NewSubscriber(sub *pulsar.PulsarClient, smtp notify.SmtpServer) Subscriber {
	return Subscriber{
		sub:  sub,
		smtp: &smtp,
	}
}

func (s Subscriber) SubscribeToVerifyEmail() {
	go func() {
		for {
			//Create event consumption
			sub, err := s.sub.CreateConsumer(constant.VERIFYEMAIL, constant.SUBSCRIPTION)
			if err != nil {
				log.Printf("error creating consumer %v", err)
				continue
			}
			msg, err := sub.ReceiveMessage()
			if err != nil {
				log.Printf("error receiving message %v", err)
				continue
			}

			//Convert the interface msg to string
			result := msg.(string)

			//Unmarshal the result into verify variable
			var verify models.VerifyEmail
			if err := json.Unmarshal([]byte(result), &verify); err != nil {
				log.Printf("error un marshaling %v", err)
				continue
			}

			//Send verify email to the customer
			data := models.VerifyEmail{
				Email: verify.Email,
				Code:  verify.Code,
			}
			if err := s.smtp.SendVerifyEmail(data); err != nil {
				log.Printf("error sending verify email %v", err)
				continue
			}
		}

	}()
}

// Shutdown gracefully closes the Pulsar client
func (s Subscriber) Shutdown() {
	//Pulsar is shutting down
	log.Println("Consumer shut down")
	s.sub.Close()
}
