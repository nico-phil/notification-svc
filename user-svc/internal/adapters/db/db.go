package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/nico-phil/notification/user-svc/internal/application/core/domain"
)

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

func(a *Adapter) SaveUser(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users(firstname, lastname, email, password)
		VALUES($1, $2, $3, $4)
		RETURNING id
	`
	args := []any{user.Firstname, user.Lastname, user.Email, user.HashPassword}
	return a.db.QueryRowContext(ctx, query, args...).Scan(&user.ID)
}

func(a *Adapter) SaveDevice(ctx context.Context, device *domain.Device) error {
	query := `
		INSERT INTO devices(device_type, device_token, user_id)
		VALUES($1, $2, $3)
		RETURNING id, user_id
	`
	args := []any{device.DeviceType, device.DeviceToken, device.UserID}
	return a.db.QueryRowContext(ctx, query, args...).Scan(&device.ID, &device.UserID)
}