package grpc

import (
	"context"

	notifs "github.com/nico-phil/notification-proto/golang/notification"
)

func(a Adapter) Send(ctx context.Context, request *notifs.SendNotificationsRequest)(*notifs.SendNotificationsResponse, error){
	// a.api.SendNotification()
	return &notifs.SendNotificationsResponse{Send: true}, nil	
}