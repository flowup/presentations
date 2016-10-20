package email

type MockEmailSender struct {
	// defines if mail was sent or not (defaults to false)
	MailSent bool
}

func (sender *MockEmailSender) SendEmail(from, to, content string) {
	// set sent mail descriptor to true
	sender.MailSent = true
}
