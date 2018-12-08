package mailers_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/stanislas-m/mybuffaloproject/mailers"
)

func Test_SendRegistrationMail(t *testing.T) {
	r := require.New(t)
	ms, v := mailers.SMTP.(*mocksmtp.MockSMTP)
	r.True(v)
	// Clear SMTP queue after test
	defer ms.Clear()
	err := mailers.SendRegistrationMail("test@gobuffalo.io")
	r.NoError(err)
	// Test email
	r.Equal(1, ms.Count())
	m, err := ms.LastMessage()
	r.NoError(err)
	r.Equal("Buffalo <noreply@gobuffalo.io>", m.From)
	r.Equal("Registration succeed", m.Subject)
	r.Equal(1, len(m.To))
	r.Equal("test@gobuffalo.io", m.To[0])
}
