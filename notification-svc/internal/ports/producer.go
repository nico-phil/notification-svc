package ports

import "github.com/nico-phil/notification/internal/application/core/domain"

type ProducerPort interface {
	PushMessageToQueue(topic string, message domain.PushNotification) error
}