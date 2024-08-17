package ports

import (
	"context"

	"github.com/nico-phil/notification/internal/application/core/domain"
)

type APIPort interface {
	SendPushNotification(context.Context, domain.PushNotification) bool
	GetDevice(context.Context, int64) (domain.Device, error)
}