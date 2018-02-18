package main

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"
)

type EmailUser struct {
	User string
	Pass string
	Host string
	Port string
}

type SmtpTemplateData struct {
	From    string
	To      string
	Subject string
	Body    string
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}

Sincerely,

{{.From}}
`

func sendMail(data *SmtpTemplateData) error {
	var doc bytes.Buffer

	t := template.New("emailTemplate")
	t, err := t.Parse(emailTemplate)
	if err != nil {
		return err
	}

	err = t.Execute(&doc, data)
	if err != nil {
		return err
	}

	emailUser := &EmailUser{os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASS"), os.Getenv("EMAIL_HOST"), os.Getenv("EMAIL_PORT")}

	auth := smtp.PlainAuth(
		"",
		emailUser.User,
		emailUser.Pass,
		emailUser.Host,
	)

	err = smtp.SendMail(emailUser.Host+":"+emailUser.Port,
		auth,
		emailUser.User,
		[]string{data.To},
		doc.Bytes())
	if err != nil {
		return err
	}

	return nil
}
