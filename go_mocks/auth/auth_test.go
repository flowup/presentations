package auth

import (
	"testing"

	"github.com/flowup/presentations/go_mocks/auth"
	"github.com/flowup/presentations/go_mocks/email"
	"github.com/stretchr/testify/assert"
)

// test valid password
func TestRegisterConfirmationEmail(t *testing.T) {
	mockEmail := &email.MockEmailSender{}
	authService := &auth.AuthService{EmailSender: mockEmail}

	authService.Register("peter.malina@flowup.cz", "12345678")
	assert.Equal(t, true, mockEmail.MailSent)
}
