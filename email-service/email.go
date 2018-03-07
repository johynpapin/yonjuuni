package main

import (
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	Name    string
	Address string
}

func sendMail(from Email, to Email, subject string, plainTextContent string, htmlContent string) error {
	message := mail.NewSingleEmail(mail.NewEmail(from.Name, from.Address), subject, mail.NewEmail(to.Name, to.Address), plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)

	return err
}
