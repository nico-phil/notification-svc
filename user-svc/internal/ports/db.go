package ports

import (
	"context"

	"github.com/nico-phil/notification/user-svc/internal/application/core/domain"
)

type DBPort interface {
	SaveUser(context.Context, *domain.User) error
	SaveDevice(context.Context, *domain.Device) error
}