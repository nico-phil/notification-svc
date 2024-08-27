package api

import (
	"context"
	"errors"

	"github.com/nico-phil/notification/internal/application/core/domain"
	"github.com/nico-phil/notification/internal/ports"
)

var userCache map[int64]domain.User = map[int64]domain.User{}
var deviceCache map[int64] domain.Device = map[int64]domain.Device{}

type Application struct {
	producer ports.ProducerPort
	user ports.UserPort
}

func NewApplication(producer ports.ProducerPort, user ports.UserPort) *Application{
	return &Application{producer: producer, user: user}
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

	var err error
	device, ok := deviceCache[notification.UserId]
	if !ok {
		device, err = a.user.GetDevice(ctx, notification.UserId)
		deviceCache[notification.UserId] = device
	}
	
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
	var err error
	u, ok := userCache[notification.UserId]
	if !ok {
		u, err = a.user.Get(ctx, notification.UserId)
		userCache[notification.UserId] = u
	}

	if err != nil {
		return err
	}

	emailNotification := domain.EmailNotification {
		Title: notification.Title,
		Content: notification.Content,
		Email: u.Email,
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
	return a.producer.PushMessageToQueueSMS("SMS_QUEUE", smsNotification)

}


