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

func NewSubscriber(sub *pulsar.PulsarClient, smtp *notify.SmtpServer) Subscriber {
	return Subscriber{
		sub:  sub,
		smtp: smtp,
	}
}

/*
Subscribe to users Produced records to pulser broker
*/
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

func (s Subscriber) SubscribeToWelcomeEmail() {
	go func() {
		for {
			//Create event consumption
			sub, err := s.sub.CreateConsumer(constant.WELCOME, constant.SUBSCRIPTION)
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
			var verify models.WelcomeMessage
			if err := json.Unmarshal([]byte(result), &verify); err != nil {
				log.Printf("error un marshaling %v", err)
				continue
			}

			//Send verify email to the customer
			data := models.WelcomeMessage{
				Email: verify.Email,
				Name:  verify.Name,
			}
			if err := s.smtp.SendWelcomeEmail(data); err != nil {
				log.Printf("error sending welcome email %v", err)
				continue
			}
		}

	}()
}

func (s Subscriber) SubscribeToSendForgotPasswordCode() {
	go func() {
		for {
			//Create event consumption
			sub, err := s.sub.CreateConsumer(constant.FORGOTPASSWORD, constant.SUBSCRIPTION)
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
			var pass models.ForgotPassword
			if err := json.Unmarshal([]byte(result), &pass); err != nil {
				log.Printf("error un marshaling %v", err)
				continue
			}

			//Send verify email to the customer
			data := models.ForgotPassword{
				Email: pass.Email,
				Name:  pass.Name,
				Code:  pass.Code,
			}
			if err := s.smtp.SendForgotPasswordCodeEmail(data); err != nil {
				log.Printf("error sending forgot password code email %v", err)
				continue
			}
		}

	}()
}

func (s Subscriber) SubscribeToSuccessfulResetPassword() {
	go func() {
		for {
			//Create event consumption
			sub, err := s.sub.CreateConsumer(constant.RESETPASSWORD, constant.SUBSCRIPTION)
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
			var pass models.ForgotPassword
			if err := json.Unmarshal([]byte(result), &pass); err != nil {
				log.Printf("error un marshaling %v", err)
				continue
			}

			//Send verify email to the customer
			data := models.ForgotPassword{
				Email: pass.Email,
				Name:  pass.Name,
			}
			if err := s.smtp.SendSuccessfulResetPasswordEmail(data); err != nil {
				log.Printf("error sending reset password email %v", err)
				continue
			}
		}

	}()
}

func (s Subscriber) SubscribeToSuccessfulResetChangePassword() {
	go func() {
		for {
			//Create event consumption
			sub, err := s.sub.CreateConsumer(constant.SUCCESSCHANGEPASSWORD, constant.SUBSCRIPTION)
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
			var pass models.ForgotPassword
			if err := json.Unmarshal([]byte(result), &pass); err != nil {
				log.Printf("error un marshaling %v", err)
				continue
			}

			//Send verify email to the customer
			data := models.ForgotPassword{
				Email: pass.Email,
				Name:  pass.Name,
			}
			if err := s.smtp.SendSuccessfulResetPasswordChangeEmail(data); err != nil {
				log.Printf("error sending reset password email %v", err)
				continue
			}
		}

	}()
}

func (s Subscriber) SubscribeToSuccessfulAidCreation() {
	go func() {
		for {
			//Create event consumption
			sub, err := s.sub.CreateConsumer(constant.AIDCREATED, constant.SUBSCRIPTION)
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
			var pass models.ForgotPassword
			if err := json.Unmarshal([]byte(result), &pass); err != nil {
				log.Printf("error un marshaling %v", err)
				continue
			}

			//Send verify email to the customer
			data := models.ForgotPassword{
				Email: pass.Email,
				Name:  pass.Name,
			}
			if err := s.smtp.SendSuccessfulMessageAidCreated(data); err != nil {
				log.Printf("error sending reset password email %v", err)
				continue
			}
		}

	}()
}

func (s Subscriber) SubscribeToSuccessfulAidUpdating() {
	go func() {
		for {
			//Create event consumption
			sub, err := s.sub.CreateConsumer(constant.AIDUPDATED, constant.SUBSCRIPTION)
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
			var pass models.ForgotPassword
			if err := json.Unmarshal([]byte(result), &pass); err != nil {
				log.Printf("error un marshaling %v", err)
				continue
			}

			//Send verify email to the customer
			data := models.ForgotPassword{
				Email: pass.Email,
				Name:  pass.Name,
			}
			if err := s.smtp.SendSuccessfulMessageAidUpdated(data); err != nil {
				log.Printf("error sending reset password email %v", err)
				continue
			}
		}

	}()
}

func (s Subscriber) SubscribeToSuccessfulAidDeleted() {
	go func() {
		for {
			//Create event consumption
			sub, err := s.sub.CreateConsumer(constant.AIDDELETED, constant.SUBSCRIPTION)
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
			var pass models.ForgotPassword
			if err := json.Unmarshal([]byte(result), &pass); err != nil {
				log.Printf("error un marshaling %v", err)
				continue
			}

			//Send verify email to the customer
			data := models.ForgotPassword{
				Email: pass.Email,
				Name:  pass.Name,
			}
			if err := s.smtp.SendSuccessfulMessageAidDeleted(data); err != nil {
				log.Printf("error sending reset password email %v", err)
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
