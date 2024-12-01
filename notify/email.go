package notify

import (
	"fmt"

	"github.com/DecodeWorms/notification.sv/models"
)

func (sm SmtpServer) SendVerifyEmail(data models.VerifyEmail) error {
	to := []string{
		data.Email,
	}
	message := []byte(data.Code)
	if err := sm.SendEmail(to, message); err != nil {
		return err
	}
	return nil
}

func (sm SmtpServer) SendWelcomeEmail(data models.WelcomeMessage) error {
	to := []string{
		data.Email,
	}
	data.Message = fmt.Sprintf("Welcome Mr/Mrs %s,  you have completed the kyc", data.Name)
	msg := data.Message
	if err := sm.SendEmail(to, []byte(msg)); err != nil {
		return err
	}
	return nil
}
