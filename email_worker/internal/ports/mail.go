package ports

type MailPort interface {
	SendRequestToMailSender() error
}