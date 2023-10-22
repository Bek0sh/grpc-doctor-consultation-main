package email

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

type GmailSender struct {
	name             string
	fromEmailAddress string
	fromEmailPass    string
}

func NewGmailSender(name string, fromEmailAddress string, fromEmailPass string) EmailSender {
	return &GmailSender{name: name, fromEmailAddress: fromEmailAddress, fromEmailPass: fromEmailPass}
}

func (g *GmailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()

	e.From = fmt.Sprintf("%s <%s>", g.name, g.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Bcc = bcc
	e.Cc = cc

	// for _, f := range attachFiles {
	// 	_, err := e.AttachFile(f)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to attach file: %s, error: %v", f, err)
	// 	}
	// }

	smtpAuth := smtp.PlainAuth("", g.fromEmailAddress, g.fromEmailPass, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}
