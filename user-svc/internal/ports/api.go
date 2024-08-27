package ports

import (
	"context"

	"github.com/nico-phil/notification/user-svc/internal/application/core/domain"
)

type APIPort interface {
	Create(context.Context, *domain.User) error
	Get(context.Context, int64) (domain.User, error)
	CreateDevice(context.Context, *domain.Device) error
	GetUserDevice(context.Context, int64) (domain.Device, error)
}