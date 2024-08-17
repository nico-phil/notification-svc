package grpc

import (
	"context"

	notif "github.com/nico-phil/notification-proto/golang/notification"
	"github.com/nico-phil/notification/internal/application/core/domain"
)

var deviceCahe map[int64]domain.Device = make(map[int64]domain.Device)

func(a Adapter) Push(ctx context.Context, request *notif.SendPushNotificationsRequest)(*notif.SendPushNotificationsResponse, error){

	var device domain.Device
	var err error
	device, ok := deviceCahe[1]
	if !ok {
		device, err = a.api.GetDevice(ctx, 1)
	}

	if err != nil {
		return &notif.SendPushNotificationsResponse{Sent: false}, err
	}
	
	deviceCahe[int64(device.ID)] = device
	
	pushNotification := domain.NewPushNotification(request.Content, request.Title,  device)
	
	a.api.SendPushNotification(ctx, pushNotification)

	
	return &notif.SendPushNotificationsResponse{Sent: true}, nil	
}

