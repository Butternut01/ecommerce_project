package service

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"user-service/internal/config"
)

type EmailService struct {
	cfg *config.Config
}

func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{cfg: cfg}
}

func (es *EmailService) SendWelcomeEmail(to, username string) error {
	tmpl := `Subject: Welcome to Our Service!
	
Hello {{.Username}},

Thank you for registering with us!

Best regards,
The Team
`

	t, err := template.New("email").Parse(tmpl)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	if err := t.Execute(&body, struct{ Username string }{Username: username}); err != nil {
		return err
	}

	auth := smtp.PlainAuth("", es.cfg.EmailFrom, es.cfg.EmailPass, es.cfg.SMTPHost)

	return smtp.SendMail(
		fmt.Sprintf("%s:%s", es.cfg.SMTPHost, es.cfg.SMTPPort),
		auth,
		es.cfg.EmailFrom,
		[]string{to},
		body.Bytes(),
	)
}