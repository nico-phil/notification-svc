package grpc

import (
	"context"

	notifs "github.com/nico-phil/notification-proto/golang/notification"
	"github.com/nico-phil/notification/internal/application/core/domain"
)

func(a Adapter) Send(ctx context.Context, request *notifs.SendNotificationsRequest)(*notifs.SendNotificationsResponse, error){
	message := domain.NewNotification("Your driver is comming","The Good seat","123","PUSH", "", "ANDROID" )
	
	a.api.SendPushNotification(ctx, message)
	return &notifs.SendNotificationsResponse{Send: true}, nil	
}