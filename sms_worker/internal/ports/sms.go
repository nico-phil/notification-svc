package ports

type SMSPORT interface {
	SendSMSNotification(body, from, to string) error
}