package db

import (
	"context"
	"database/sql"
	"errors"

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
		INSERT INTO users(firstname, lastname, email, phone_number, password)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id
	`
	args := []any{user.Firstname, user.Lastname, user.Email, user.PhoneNumber, user.HashPassword}
	return a.db.QueryRowContext(ctx, query, args...).Scan(&user.ID)
}

func(a *Adapter) GetUser(ctx context.Context, id int64)(domain.User, error){
	var user domain.User
	query := `
		SELECT * FROM users
		WHERE id=$1
		`
	args := []any{id}
	err := a.db.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
	)
	if err!= nil {
		if errors.Is(err, sql.ErrNoRows){
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
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

func(a *Adapter) GetUserDevice(ctx context.Context, userId int64)(domain.Device, error){
	var device domain.Device
	query := `
		SELECT * FROM devices
		WHERE user_id=$1
		`
	args := []any{userId}
	err := a.db.QueryRowContext(ctx, query, args...).Scan(
		&device.ID,
		&device.DeviceToken,
		&device.DeviceType,
		&device.UserID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Device{}, errors.New("device not found")
		}
		return domain.Device{}, err
	}
	return device, nil
}	



