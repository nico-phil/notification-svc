package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
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

type Adapter struct {
	db *sql.DB
}

func NewAdapter(dsn string) (*Adapter, error){
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &Adapter{db: db}, nil
}