package notify

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/DecodeWorms/notification.sv/config"
	"github.com/DecodeWorms/notification.sv/constant"
)

type SmtpServer struct {
	Host   string
	Port   string
	config config.Config
}

func NewSmtServer(c config.Config) SmtpServer {
	return SmtpServer{
		Host:   c.Host,
		Port:   c.SmtpPort,
		config: c,
	}
}

func (sm SmtpServer) Address() string {
	if sm.Host == "" || sm.Port == "" {
		log.Println("SMTP Host or Port is missing")
		return ""
	}
	return fmt.Sprintf("%s:%s", strings.TrimSpace(sm.Host), strings.TrimSpace(sm.Port))

}

func (sm SmtpServer) SendEmail(to []string, message []byte) error {
	auth := smtp.PlainAuth("", constant.From, sm.config.Password, sm.Host)
	if err := smtp.SendMail(sm.Address(), auth, constant.From, to, message); err != nil {
		return err
	}
	return nil
}
