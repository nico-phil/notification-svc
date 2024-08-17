package ports

import (
	"context"

	"github.com/nico-phil/notification/internal/application/core/domain"
)

type DBPort interface {
	Get(context.Context, int64) (domain.Device, error)
}