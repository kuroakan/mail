package service

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"mail.project/bootstrap"
	"mail.project/entity"
	"net/smtp"
)

type MailService struct {
	cfg   *bootstrap.Config
	kafka *kafka.Reader
}

func New(cfg *bootstrap.Config) *MailService {
	return &MailService{
		cfg: cfg,
	}
}

func (ms *MailService) SendMail(mail entity.Mail) error {
	auth := smtp.PlainAuth("", ms.cfg.GmailLogin, ms.cfg.GmailPassword, ms.cfg.GmailHost)

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", mail.Receiver, mail.Subject, mail.Message))

	err := smtp.SendMail(fmt.Sprintf("%s:%s", ms.cfg.GmailHost, ms.cfg.GmailPort), auth, ms.cfg.GmailLogin, []string{mail.Receiver}, msg)
	if err != nil {
		return err
	}

	return nil
}
