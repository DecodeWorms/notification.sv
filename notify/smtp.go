package notify

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/DecodeWorms/notification.sv/config"
	"github.com/DecodeWorms/notification.sv/constant"
)

type SmtpServer struct {
	Host   string
	Port   string
	config config.Config
	Client *smtp.Client
}

func NewSmtServer(c config.Config) (*SmtpServer, error) {
	add := fmt.Sprintf("%s:%s", strings.TrimSpace(c.Host), strings.TrimSpace(c.SmtpPort))
	conn, err := smtp.Dial(add)
	if err != nil {
		return nil, fmt.Errorf("error connecting to smtp server %v", err)
	}
	if err = conn.StartTLS(&tls.Config{ServerName: c.Host}); err != nil {
		return nil, fmt.Errorf("error sending STARTTLS command %v", err)
	}
	if err = conn.Auth(smtp.PlainAuth("", constant.From, c.Password, c.Host)); err != nil {
		return nil, fmt.Errorf("error authenticating a user %v", err)
	}

	return &SmtpServer{
		Host:   c.Host,
		Port:   c.SmtpPort,
		config: c,
		Client: conn,
	}, nil
}

func (sm SmtpServer) SendEmail(to []string, message []byte) error {
	//Register the sender email address
	if err := sm.Client.Mail(constant.From); err != nil {
		return fmt.Errorf("error sending a mail command to the server %v", err)
	}
	//Process the receiver email address(es)
	for _, v := range to {
		if err := sm.Client.Rcpt(v); err != nil {
			return fmt.Errorf("error sending RCPT command %v", err)
		}
	}
	//Getting data writer
	writer, err := sm.Client.Data()
	if err != nil {
		return fmt.Errorf("error getting DATA writer: %w", err)
	}
	if _, err = writer.Write(message); err != nil {
		return fmt.Errorf("error writing email content: %w", err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("error closing DATA writer: %w", err)
	}
	return nil
}
