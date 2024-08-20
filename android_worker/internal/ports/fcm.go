package ports

type FCMPort interface {
	SendNotification(title, body, token string) error
}