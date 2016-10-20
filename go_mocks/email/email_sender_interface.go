package email

type Sender interface {
	SendEmail(from, to, content string)
}
