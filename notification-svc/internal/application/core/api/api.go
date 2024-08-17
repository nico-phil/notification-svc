package api

import (
	"context"
	"log"

	"github.com/nico-phil/notification/internal/application/core/domain"
	"github.com/nico-phil/notification/internal/ports"
)

type Application struct {
	producer ports.ProducerPort
	db ports.DBPort
}

func NewApplication(kafkaAdapter ports.ProducerPort, db ports.DBPort) *Application{
	return &Application{producer: kafkaAdapter, db: db}
}

func(a *Application) SendPushNotification(ctx context.Context, notification domain.PushNotification) bool {
	var topic string

	switch  {
	case notification.Device.DeviceType == "IOS":
		topic= "IOS_QUEUE"
	case notification.Device.DeviceType == "ANDROID":
		topic="ANDROID_QUEUE"
	// case notifType == "SMS": 
	// 	topic = "SMS_QUEUE"
	// case notifType == "EMAIL":
	// 	topic = "EMAIL_QUEUE"
	default: 
		log.Printf("unknown device type: %s", notification.Device.DeviceType)
	
	}
	
	a.producer.PushMessageToQueue(topic, notification)
	return true
}

func(a *Application) GetDevice(ctx context.Context, id int64) (domain.Device, error){
	device, err := a.db.Get(ctx, id)
	if err != nil {
		return domain.Device{}, err
	}

	return device, nil
}

