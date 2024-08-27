package domain

import "encoding/json"


type Notification struct {
	Title string `json:"title"`
	Content string `json:"content"`
	UserId int64 `json:"user_id"`
	NotificationType string `json:"notification_type"`
}

type Device struct{
	ID int64	`json:"id"`
	DeviceToken string `json:"device_token"`
	DeviceType string 	`json:"device_type"`
	UserId int64 `json:"user_id"`
}

type PushNotification struct {
	Notification Notification 	`json:"notification"`
	Device Device `json:"device"`
}

type EmailNotification struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Email string `json:"email"`
}

type SMSNotification struct {
	Title string `json:"title"`
	Content string `json:"content"`
	PhoneNumber string `json:"phone_number"`
}

type User struct {
	Firstname string `json:"first_name"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
}



func NewNotification(userId int64, title, content, notificationType string ) Notification{
	return Notification{
		Title: title,
		Content: content,
		UserId: userId,
		NotificationType: notificationType,
	}
}

func NewPushNotification(title, content string, device Device) PushNotification {
	return PushNotification {
		Notification: Notification{Content: content, Title: title},
		Device: device,
	}
}

func(m PushNotification) Encode()([]byte, error) {
	return json.Marshal(m)
}

func (m PushNotification) Length() int{
	encode, _ := m.Encode()
	return len(encode)
}

func(m EmailNotification) Encode()([]byte, error) {
	return json.Marshal(m)
}

func (m EmailNotification) Length() int{
	encode, _ := m.Encode()
	return len(encode)
}

func(m SMSNotification) Encode()([]byte, error) {
	return json.Marshal(m)
}

func (m SMSNotification) Length() int{
	encode, _ := m.Encode()
	return len(encode)
}