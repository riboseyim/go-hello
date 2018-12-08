package main

import (
	"bytes"
	"html/template"
	"log"
)

//Request struct
type MailRequest struct {
	to      string
	subject string
	body    string
}

func NewMailRequest(to string, subject string) *MailRequest {
	return &MailRequest{
		to:      to,
		subject: subject,
	}
}

func (r *MailRequest) ParseTemplate(templateFileName string, data interface{}) error {
	log.Println("-------\n ParseTemplate()  \n--------")
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()

	log.Println("-------\n ParseTemplate() r.body:%s \n--------", r.body)
	return nil
}

func (r *MailRequest) SendEmail() {

	//mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	// subject := "Subject: " + r.subject + "!\n"
	// msg := []byte(subject + mime + "\n" + r.body)

	TTankMail(r.to, r.subject, r.body)

}
