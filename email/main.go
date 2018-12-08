package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
)

var auth smtp.Auth

// my $smtp_server = "139.129.160.101";
// my $smtp_user = 'itereport@iteview.com';
// my $smtp_password = "R2e0p!O^";

const (
	Sender       = "zijingshanke@qq.com"
	SenderPass   = "vfmwxurvzqyqbdgi"
	SenderServer = "smtp.qq.com"
	SenderPort   = "smtp.qq.com:587"
	Recivers     = "yanrui@zhong-ying.com"
)

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(from string, to []string, subject, body string) *Request {
	return &Request{
		from:    from,
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	log.Println("-------\n SendEmail() \n--------")

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)

	if err := smtp.SendMail(SenderPort, auth, r.from, r.to, msg); err != nil {
		log.Println("-----sendmail() error:%s", err)
		return false, err
	}
	return true, nil
}

func main() {
	auth = smtp.PlainAuth("", Sender, SenderPass, SenderServer)
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "Dhanush",
		URL:  "http://geektrust.in",
	}

	r := NewRequest(Sender, []string{Recivers}, "Hello Rui!", "Hello, World! hahahah")
	err := r.ParseTemplate("template.html", templateData)
	if err != nil {
		ok, _ := r.SendEmail()
		fmt.Println(ok)
	}
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	log.Println("-------\n ParseTemplate() \n--------")
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()

	log.Println("-------\n %s \n--------", r.body)
	return nil
}
