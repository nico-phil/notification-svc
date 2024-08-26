package user

import (
	"github.com/nico-phil/notification-proto/golang/user"
	"google.golang.org/grpc"
)

type Adapter struct {
	userClient user.UserClient
}

func NewAdapter(userServiceUrl string) (*Adapter, error){
	var opts [] grpc.DialOption

	conn, err := grpc.NewClient(userServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := user.NewUserClient(conn)

	return &Adapter{userClient: client}, nil
}