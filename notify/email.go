package notify

import "github.com/DecodeWorms/notification.sv/models"

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
