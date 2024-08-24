package ports

type MailPort interface {
	SendRequestToMailSender(subject, text, emailTo string) error
}