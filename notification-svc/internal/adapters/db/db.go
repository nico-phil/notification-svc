package db

import (
	"context"

	"github.com/nico-phil/notification/internal/application/core/domain"
)

type DeviceEntity struct {
	DeviceToken string
	DeviceType string
}

type DBModel struct {
	Devices []DeviceEntity
}
func(d DBModel ) Get(ctx context.Context,  deviceId int64) (domain.Device , error){
	return domain.Device {
		ID: 1,
		DeviceToken: "",
	}, nil
}