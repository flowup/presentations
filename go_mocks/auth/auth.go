package auth

import "github.com/flowup/presentations/go_mocks/email"

type AuthService struct {
	EmailSender email.Sender
}

func (auth *AuthService) Register(email, password string) {
	// validate password
	if len(password) < 8 {
		return
	}

	// create user somewhere in persistent storage

	// send confirmation email
	auth.EmailSender.SendEmail("my@email.com", email, "Welcome!")
}
