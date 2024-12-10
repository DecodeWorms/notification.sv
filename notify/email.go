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
	data.Message = fmt.Sprintf("Welcome  %s,  you have completed the kyc", data.Name)
	msg := data.Message
	if err := sm.SendEmail(to, []byte(msg)); err != nil {
		return err
	}
	return nil
}

func (sm SmtpServer) SendForgotPasswordCodeEmail(data models.ForgotPassword) error {
	to := []string{
		data.Email,
	}
	msg := fmt.Sprintf("Welcome %s, we received your password change request and here is your confirmation code %s", data.Name, data.Code)
	if err := sm.SendEmail(to, []byte(msg)); err != nil {
		return err
	}
	return nil
}

func (sm SmtpServer) SendSuccessfulResetPasswordEmail(data models.ForgotPassword) error {
	to := []string{
		data.Email,
	}
	msg := fmt.Sprintf("Hi %s, password reset was successful", data.Name)
	if err := sm.SendEmail(to, []byte(msg)); err != nil {
		return err
	}
	return nil
}

func (sm SmtpServer) SendSuccessfulResetPasswordChangeEmail(data models.ForgotPassword) error {
	to := []string{
		data.Email,
	}
	msg := fmt.Sprintf("Hi %s, password change was successful", data.Name)
	if err := sm.SendEmail(to, []byte(msg)); err != nil {
		return err
	}
	return nil
}
