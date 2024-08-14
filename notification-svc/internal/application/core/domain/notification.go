package domain

import "encoding/json"


type Notification struct {
	Content string `json:"content"`
	From string `json:"from"`
	To string `json:"to"`
	NotifType string `json:"notif_type"`
	Device Device `json:"device"`

}

type Device struct{
	DeviceToken string `json:"device_token"`
	DeviceType string 	`json:"device_type"`
}

func NewNotification(content, from, to, notifType, deviceToken, deviceType string) Notification {
	return Notification {
		Content: content,
		From: from,
		To: to,
		NotifType: "PUSH",
		Device: Device{DeviceToken: deviceToken, DeviceType: deviceType},
	}
}

func(n Notification) Encode()([]byte, error) {
	return json.Marshal(n)
}

func (n Notification) Length() int{
	encode, _ := n.Encode()
	return len(encode)
}