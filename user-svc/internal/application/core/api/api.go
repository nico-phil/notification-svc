package api

import (
	"context"
	"log"

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

	err := user.EncriptPassword()
	if err != nil{
		log.Println("error hasing password", err)
	}
	
	err = a.db.SaveUser(ctx, user)
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

func(a *Application) GetUserDevice(ctx context.Context, userId int64)(domain.Device, error){
	device, err := a.db.GetUserDevice(ctx, userId)
	if err != nil {
		return domain.Device{}, err
	}

	return device, nil
}
