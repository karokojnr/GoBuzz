package mailer

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendgridMailer struct {
	fromEmail string
	apiKey    string
	client    *sendgrid.Client
}

func NewSendgrid(fromEmail, apiKey string) *SendgridMailer {
	client := sendgrid.NewSendClient(apiKey)
	return &SendgridMailer{
		fromEmail: fromEmail,
		apiKey:    apiKey,
		client:    client,
	}
}

func (s *SendgridMailer) Send(templateFile, username, email string, data any, isSandbox bool) error {
	from := mail.NewEmail(FromName, s.fromEmail)
	to := mail.NewEmail(username, email)

	subject := new(bytes.Buffer)
	body := new(bytes.Buffer)

	message := mail.NewSingleEmail(from, subject.String(), to, "", body.String())
	message.SetMailSettings(&mail.MailSettings{
		SandboxMode: &mail.Setting{
			Enable: &isSandbox,
		},
	})

	for i := 0; i < maxRetries; i++ {
		res, err := s.client.Send(message)
		if err != nil {
			log.Printf("Failed to send email to %v, attempt %d of %d \n", email, i+1, maxRetries)
			log.Printf("Error: %v \n", err)

			// exponential backoff
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}

		log.Printf("Email sent with status code: %v \n", res.StatusCode)
		return nil
	}
	return fmt.Errorf("Failed to send email to %v after %d attempts", email, maxRetries)
}
