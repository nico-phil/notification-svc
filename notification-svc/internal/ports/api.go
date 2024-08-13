package ports

import (
	"context"

	"github.com/nico-phil/notification/internal/application/core/domain"
)

type APIPort interface {
	SendNotification(context.Context, domain.Notification) bool
}