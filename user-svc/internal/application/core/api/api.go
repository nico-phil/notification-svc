package api

import (
	"context"

	"github.com/nico-phil/notification/user-svc/internal/application/core/domain"
	"github.com/nico-phil/notification/user-svc/internal/ports"
)


type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application{
	return &Application{
		db: db,
	}
}

func(a *Application) CreateUser(ctx context.Context,  user *domain.User) error {
	err := a.db.SaveUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func(a *Application) CreateDevice(ctx context.Context, device *domain.Device) error {
	err := a.db.SaveDevice(ctx, device)
	if err != nil {
		return err
	}
	return nil
}