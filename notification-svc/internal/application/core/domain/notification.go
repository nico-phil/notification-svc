package domain

import "encoding/json"


type Notification struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

type Device struct{
	ID int	`json:"id"`
	DeviceToken string `json:"device_token"`
	DeviceType string 	`json:"device_type"`
}

type PushNotification struct {
	Notification Notification 	`json:"notification"`
	Device Device `json:"device"`
}

func NewPushNotification(content, title string, device Device) PushNotification {
	return PushNotification {
		Notification: Notification{Content: content, Title: title},
		Device: device,
	}
}

func(n PushNotification) Encode()([]byte, error) {
	return json.Marshal(n)
}

func (n PushNotification) Length() int{
	encode, _ := n.Encode()
	return len(encode)
}