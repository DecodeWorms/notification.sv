package notify

import (
	"net/smtp"

	"github.com/DecodeWorms/notification.sv/config"
	"github.com/DecodeWorms/notification.sv/constant"
)

type SmtpServer struct {
	Host string
	Port string
}

func NewSmtServer(h, p string) SmtpServer {
	return SmtpServer{
		Host: h,
		Port: p,
	}
}

func (sm SmtpServer) Address() string {
	return sm.Host + sm.Port
}

func (sm SmtpServer) SendEmail(to []string, message []byte) error {
	c := config.ImportConfig(config.OSSource{})
	smtpServer := SmtpServer{Host: c.Host, Port: c.Port}
	auth := smtp.PlainAuth("", constant.From, c.Password, sm.Host)
	if err := smtp.SendMail(smtpServer.Address(), auth, constant.From, to, message); err != nil {
		return err
	}
	return nil
}
