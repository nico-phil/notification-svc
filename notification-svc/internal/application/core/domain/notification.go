package domain


type Notification struct {
	Text string `json:"text"`
}

func NewNotification(text string) Notification {
	return Notification{Text: text}
}