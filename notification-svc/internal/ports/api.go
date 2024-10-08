package ports

import (
	"context"

	"github.com/nico-phil/notification/internal/application/core/domain"
)

type APIPort interface {
	SendPushNotification(context.Context, domain.Notification) error
	SendNotification(context.Context, domain.Notification) error

}