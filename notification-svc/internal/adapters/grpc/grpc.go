package grpc

import (
	"context"

	notif "github.com/nico-phil/notification-proto/golang/notification"
	"github.com/nico-phil/notification/internal/application/core/domain"
)

func(a Adapter) Push(ctx context.Context, request *notif.SendPushNotificationsRequest)(*notif.SendPushNotificationsResponse, error){

	device, err := a.api.GetDevice(ctx, 1)
	if err != nil {
		return &notif.SendPushNotificationsResponse{Sent: false}, err
	}
	pushNotification := domain.NewPushNotification("Hello Friend", "Gretting",  device)
	
	a.api.SendPushNotification(ctx, pushNotification)

	
	return &notif.SendPushNotificationsResponse{Sent: true}, nil	
}

