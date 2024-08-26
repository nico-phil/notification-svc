package ports

import (
	"context"

	"github.com/nico-phil/notification/internal/application/core/domain"
)

type UserPort interface{
	GetDevice(context.Context,  int64) (domain.Device, error)
}