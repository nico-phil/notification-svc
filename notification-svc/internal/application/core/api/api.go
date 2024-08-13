package api

import (
	"context"

	"github.com/nico-phil/notification/internal/application/core/domain"
)

type Application struct {

}

func NewApplication() *Application{
	return &Application{}
}

func(a *Application) SendNotification(ctx context.Context, notification domain.Notification) bool {
	return true
}


