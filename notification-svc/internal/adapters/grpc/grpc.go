package grpc

import (
	"context"

	notifs "github.com/nico-phil/notification-proto/golang/notification"
	"github.com/nico-phil/notification/internal/application/core/domain"
)

func(a Adapter) Send(ctx context.Context, request *notifs.SendNotificationsRequest)(*notifs.SendNotificationsResponse, error){

	device, err := a.api.GetDevice(ctx, 1)
	if err != nil {
		return &notifs.SendNotificationsResponse{Send: false}, err
	}
	pushNotification := domain.NewPushNotification("Hello Friend", "Gretting",  device)
	
	a.api.SendPushNotification(ctx, pushNotification)

	
	return &notifs.SendNotificationsResponse{Send: true}, nil	
}

