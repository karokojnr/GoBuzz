package mailer

import "github.com/sendgrid/sendgrid-go"

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
