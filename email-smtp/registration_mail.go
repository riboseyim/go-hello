package mailers

import (
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/pkg/errors"
)

func SendRegistrationMail(email string) error {
	m := mail.NewMessage()

	m.Subject = "Registration succeed"
	m.From = "Buffalo <noreply@gobuffalo.io>"
	m.To = []string{email}
	err := m.AddBody(r.HTML("registration_mail.html"), render.Data{})
	if err != nil {
		return errors.WithStack(err)
	}
	if err := SMTP.Send(m); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
