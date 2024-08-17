package db

import (
	"context"
	"database/sql"
	"fmt"

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

func(a *Adapter) Save(ctx context.Context, device *domain.Device) error {
	query := `
		INSERT INTO devices(deviceToken, deviceType)
		VALUES($1, $2)
		RETURNING id
	`
	args := []any{device.DeviceToken, device.DeviceType}
	fmt.Println(device)
	return a.db.QueryRowContext(ctx, query, args...).Scan(&device.ID)
}

func(a *Adapter) Get(ctx context.Context, id int64)(domain.Device, error){
	var device domain.Device
	query := `
		SELECT * FROM devices
		WHERE id=$1
		`
	args := []any{id}
	err := a.db.QueryRowContext(ctx, query, args...).Scan(
		&device.ID,
		&device.DeviceToken,
		&device.DeviceType,
	)
	if err != nil {
		return domain.Device{}, err
	}
	return device, nil
}	