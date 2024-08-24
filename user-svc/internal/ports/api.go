package ports

import (
	"context"

	"github.com/nico-phil/notification/user-svc/internal/application/core/domain"
)

type APIPort interface {
	CreateUser(context.Context, domain.User)
	CreateDevice(context.Context, domain.Device)
}