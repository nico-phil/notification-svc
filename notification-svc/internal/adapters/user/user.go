package user

import (
	"context"

	"github.com/nico-phil/notification-proto/golang/user"
	"github.com/nico-phil/notification/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	userClient user.UserClient
}

func NewAdapter(userServiceUrl string) (*Adapter, error){
	var opts [] grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(userServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := user.NewUserClient(conn)

	return &Adapter{userClient: client}, nil
}

func(a *Adapter) GetDevice(ctx context.Context, userId int64)(domain.Device, error){
	r, err := a.userClient.GetDevice(ctx, &user.GetUserDeviceRequest{UserId: userId})
	if err != nil {
		return domain.Device{}, err
	}

	device := domain.Device{
		ID: r.Id,
		UserId: r.UserId,
		DeviceToken: r.DeviceToken,
		DeviceType: r.DeviceType,
	}

	return device, nil
}

func(a *Adapter) Get(ctx context.Context, id int64) (domain.User, error){
	r, err := a.userClient.Get(ctx, &user.GetUserRequest{UserId: id})
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		Firstname: r.Firstname,
		Lastname: r.Lastname,
		Email: r.Email,
		PhoneNumber: "",
	}, nil
}