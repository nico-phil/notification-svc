package api

import (
	"context"
	"log"

	"github.com/nico-phil/notification/internal/application/core/domain"
	"github.com/nico-phil/notification/internal/ports"
)

type Application struct {
	producer ports.ProducerPort
}

func NewApplication(kafkaAdapter ports.ProducerPort) *Application{
	return &Application{producer: kafkaAdapter}
}

func(a *Application) SendPushNotification(ctx context.Context, notification domain.Notification) bool {
	
	deviceType := notification.Device.DeviceType
	var topic string

	switch  {
	case deviceType == "IOS":
		topic= "IOS_QUEUE"
	case deviceType == "ANDROID":
		topic="ANDROID_QUEUE"
	// case notifType == "SMS": 
	// 	topic = "SMS_QUEUE"
	// case notifType == "EMAIL":
	// 	topic = "EMAIL_QUEUE"
	default: 
		log.Printf("unknown device or type: %s, %s", deviceType, notification.Device.DeviceToken)
	
	}
	
	a.producer.PushMessageToQueue(topic, notification)
	return true
}


