package email

import "fmt"

// GMailSender implements email.Sender that sends mails via gmail
type GMailSender struct{}

func (gmail *GMailSender) SendEmail(from, to, content string) {
	fmt.Println("Email was sent by GMail!")
}

// OutlookMailSender implements email.Sender that sends mails via outlook
type OutlookMailSender struct{}

func (outlook *OutlookMailSender) SendEmail(from, to, content string) {
	fmt.Println("Email was sent by Outlook!")
}
