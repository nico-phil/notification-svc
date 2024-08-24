package grpc

import (
	"context"
	"fmt"

	"github.com/nico-phil/notification-proto/golang/user"
)

func(a *Adapter)Create(ctx context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	fmt.Println("Hello")
	return nil, nil

}