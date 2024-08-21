package grpc

import (
	"context"

	"github.com/nico-phil/notification-proto/golang/notification/v2"
	"github.com/nico-phil/notification/internal/application/core/domain"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(a Adapter) Send(ctx context.Context, request *notification.SendNotificationRequest)(*notification.SendNotificationResponse, error){

	var validationErrors []*errdetails.BadRequest_FieldViolation
	if  len(request.Title) == 0 {
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

	if request.UserId < 1 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field: "title",
			Description: "device id cannot be less than 1",
		})
	}

	if len(validationErrors) > 0 {
		stat := status.New(codes.InvalidArgument , "invalid push notification request" )
		badRequest := &errdetails.BadRequest{}
		badRequest.FieldViolations = validationErrors
		s, _ := stat.WithDetails(badRequest)
		return nil, s.Err()
	}

	newNotification := domain.NewNotification(request.UserId, request.Title, request.Content, request.NotificationType)
	
	err := a.api.SendNotification(ctx, newNotification)
	if err != nil {
		return nil, err
	}

	
	return &notification.SendNotificationResponse{Sent:true }, nil	
}

func (a Adapter) Email() {

}

