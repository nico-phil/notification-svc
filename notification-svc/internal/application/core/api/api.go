package api

import (
	"context"
	"errors"

	"github.com/nico-phil/notification/internal/application/core/domain"
	"github.com/nico-phil/notification/internal/ports"
)

type Application struct {
	producer ports.ProducerPort
	db ports.DBPort
}

func NewApplication(producer ports.ProducerPort, db ports.DBPort) *Application{
	return &Application{producer: producer, db: db}
}

func(a *Application) SendNotification(ctx context.Context, notification domain.Notification) error{

	switch  {
	case notification.NotificationType == "PUSH":
		return a.SendPushNotification(ctx, notification)
	case notification.NotificationType == "EMAIL":
		return a.SendEmailNotification(ctx, notification)
	case notification.NotificationType == "SMS": 

	default: 
		return errors.New("unknown notification type")
	
	}

	return nil
	
}

func(a *Application) SendPushNotification(ctx context.Context, notification domain.Notification) error {
	device, err := a.db.Get(ctx, notification.UserId)
	if err != nil {
		return err
	}

	var topic string

	if device.DeviceType == "ANDROID" {
		topic = "ANDROID_QUEUE"
	}else{
		topic = "IOS_QUEUE"
	}

	newPushNotification := domain.NewPushNotification(notification.Title, notification.Content, device)

	return a.producer.PushMessageToQueue(topic, newPushNotification)
}

func(a *Application) SendEmailNotification(ctx context.Context, notification domain.Notification) error {
	email := "yyyeenphilibert17@gmail.com"

	emailNotification := domain.EmailNotification {
		Title: notification.Title,
		Content: notification.Content,
		Email: email,
	}
	return a.producer.PushMessageToQueueEmail("EMAIL_QUEUE", emailNotification)


}

func(a *Application) SendSMSNotification(ctx context.Context, notification domain.Notification) error {
	phoneNumber := "+1234566778"

	smsNotification := domain.SMSNotification {
		Title: notification.Title,
		Content: notification.Content,
		PhoneNumber: phoneNumber,
	}
	return a.producer.PushMessageToQueueSMS("EMAIL_QUEUE", smsNotification)


}


func(a *Application) GetDevice(ctx context.Context, id int64) (domain.Device, error){
	device, err := a.db.Get(ctx, id)
	if err != nil {
		return device, err
	}

	return device, nil
}

func(a *Application) SaveDevice(ctx context.Context, device *domain.Device) error {
	return a.db.Save(ctx, device)
}


