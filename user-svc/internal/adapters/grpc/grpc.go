package grpc

import (
	"context"

	"github.com/nico-phil/notification-proto/golang/user"
	"github.com/nico-phil/notification/user-svc/internal/application/core/domain"
)

func(a *Adapter)Create(ctx context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	newUser := domain.User {
		Firstname: request.Firstname,
		Lastname: request.Lastname,
		Email: request.Email,
		Password:  request.Password,
	}
	err := a.api.CreateUser(ctx, &newUser)
	if err != nil {
		return  nil, err
	}
	return &user.CreateUserResponse{Firstname: newUser.Firstname, Lastname: newUser.Lastname, Email: newUser.Email }, nil
}

func(a *Adapter)CreateDevice(ctx context.Context, request *user.CreateDeviceRequest) (*user.CreateDeviceResponse, error){

	newDevice := domain.Device {
		DeviceToken: request.DeviceToken,
		DeviceType: request.DeviceType,
		UserID: request.UserId,
	}

	err := a.api.CreateDevice(ctx, &newDevice)
	if err != nil {
		return nil, err
	}

	return &user.CreateDeviceResponse{Id: newDevice.ID}, nil
}

func(a *Adapter) GetDevice(ctx context.Context, request *user.GetUserDeviceRequest)(*user.GetUserDeviceResponse, error){
	device, err := a.api.GetUserDevice(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	return &user.GetUserDeviceResponse{Id: device.ID ,DeviceToken: device.DeviceToken, DeviceType: device.DeviceType, UserId: device.UserID}, nil
}