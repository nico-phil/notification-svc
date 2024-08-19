package grpc

import (
	"context"

	notif "github.com/nico-phil/notification-proto/golang/notification"
	"github.com/nico-phil/notification/internal/application/core/domain"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var deviceCahe map[int64]domain.Device = make(map[int64]domain.Device)

func(a Adapter) Push(ctx context.Context, request *notif.SendPushNotificationsRequest)(*notif.SendPushNotificationsResponse, error){

	var validationErrors []*errdetails.BadRequest_FieldViolation
	if len(request.Title) == 0 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field: "title",
			Description: "title cannot be empty",
		})
	}

	if len(request.Content) == 0 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field: "content",
			Description: "content cannot be empty",
		})
	}

	if len(request.DeviceId) == 0 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field: "title",
			Description: "device id cannot be negative",
		})
	}

	if len(validationErrors) > 0 {
		stat := status.New(codes.InvalidArgument , "invalid push notification request" )
		badRequest := &errdetails.BadRequest{}
		badRequest.FieldViolations = validationErrors
		s, _ := stat.WithDetails(badRequest)
		return nil, s.Err()
	}

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
	
	r := a.api.SendPushNotification(ctx, pushNotification)

	
	return &notif.SendPushNotificationsResponse{Sent: r}, nil	
}

