package ports

import (
	"context"

	"github.com/nico-phil/notification/internal/application/core/domain"
)

type UserPort interface{
	Get(context.Context, int64)(domain.User, error)
	GetDevice(context.Context,  int64) (domain.Device, error)
}