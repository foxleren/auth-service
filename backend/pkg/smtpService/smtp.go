package smtpService

import (
	"fmt"
	"net/smtp"
)

type SmtpService interface {
	SendEmail(email *EmailData) error
}

type SMTPConfig struct {
	SenderEmail    string
	SenderPassword string
	Host           string
	Port           string
}

type SmtpProvider struct {
	config *SMTPConfig
}

func NewSmtpProvider(config *SMTPConfig) *SmtpProvider {
	return &SmtpProvider{config: config}
}

func (p *SmtpProvider) SendEmail(emailData *EmailData) error {

	from := p.config.SenderEmail
	to := emailData.Recipient

	password := p.config.SenderPassword

	auth := smtp.PlainAuth("", from, password, p.config.Host)

	email := buildEmail(to, emailData)

	if err := smtp.SendMail(p.config.Host+":"+p.config.Port, auth, from, []string{to}, []byte(email)); err != nil {
		return err
	}

	return nil
}

func buildEmail(to string, emailData *EmailData) string {
	return fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", to, emailData.Subject, emailData.Content)
}
