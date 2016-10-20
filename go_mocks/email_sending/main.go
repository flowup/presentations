package main

import "github.com/flowup/presentations/go_mocks/email"

type AuthService struct {
	emailSender email.Sender
}

func (auth *AuthService) OnUserRegister(email string) {
	auth.emailSender.SendEmail("my@email.com", email, "Welcome!")
}

func main() {
	authServiceWithGMail := &AuthService{&email.GMailSender{}}
	authServiceWithOutlook := &AuthService{&email.OutlookMailSender{}}

	// this should be called after some user registered
	authServiceWithGMail.OnUserRegister("peter.malina@flowup.cz")
	authServiceWithOutlook.OnUserRegister("peter.malina@flowup.cz")
}
