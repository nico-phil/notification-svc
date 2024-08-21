package ports

import "github.com/nico-phil/notification/internal/application/core/domain"

type ProducerPort interface {
	PushMessageToQueue(topic string, message domain.PushNotification) error
	PushMessageToQueueEmail(topic string, message domain.EmailNotification) error
	PushMessageToQueueSMS(topic string, message domain.SMSNotification) error
	
	
}