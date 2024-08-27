package grpc

import (
	"context"

	"github.com/nico-phil/notification-proto/golang/user"
	"github.com/nico-phil/notification/user-svc/internal/application/core/domain"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func(a *Adapter)Create(ctx context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	
	//check input
	var validationErrors []*errdetails.BadRequest_FieldViolation
	if len(request.Email) == 0 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field: "email",
			Description: "email should not be empty",
		})
	}

	if len(request.Firstname) == 0 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field: "firstname",
			Description: "firstname should not be empty",
		})
	}

	if len(request.Lastname) == 0 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field: "lastname",
			Description: "lastname should not be empty",
		})
	}

	if len(request.Password) == 0 {
		validationErrors = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field: "password",
			Description: "password should not be empty",
		})
	}

	if len(validationErrors) > 0 {
		st := status.New(codes.InvalidArgument, "invalid request")
		badRequest := errdetails.BadRequest{}
		badRequest.FieldViolations = validationErrors
		st,_ = st.WithDetails(&badRequest)
		return nil, st.Err()
	}
	
	
	newUser := domain.User {
		Firstname: request.Firstname,
		Lastname: request.Lastname,
		Email: request.Email,
		Password:  request.Password,
	}
	err := a.api.Create(ctx, &newUser)
	if err != nil {
		return  nil, err
	}
	return &user.CreateUserResponse{Firstname: newUser.Firstname, Lastname: newUser.Lastname, Email: newUser.Email }, nil
}

func(a *Adapter) Get(ctx context.Context, request *user.GetUserRequest) (*user.GetUserResponse, error){
	u, err := a.api.Get(ctx, request.UserId)
	if err != nil {
		if err.Error() == "user not found"{
			st := status.New(codes.NotFound, "user not found")
			return nil, st.Err() 
		}
		return nil, err
	}
	return &user.GetUserResponse{
		Firstname: u.Firstname,
		Lastname: u.Lastname,
		Email: u.Email,
	}, err
}

func(a *Adapter)CreateDevice(ctx context.Context, request *user.CreateDeviceRequest) (*user.CreateDeviceResponse, error){

	newDevice := domain.Device {
		DeviceToken: request.DeviceToken,
		DeviceType: request.DeviceType,
		UserID: request.UserId,
	}

	err := a.api.CreateDevice(ctx, &newDevice)
	if err != nil {
		return nil, err
	}

	return &user.CreateDeviceResponse{Id: newDevice.ID}, nil
}

func(a *Adapter) GetDevice(ctx context.Context, request *user.GetUserDeviceRequest)(*user.GetUserDeviceResponse, error){
	device, err := a.api.GetUserDevice(ctx, request.UserId)
	if err != nil {
		if err.Error() == "device not found" {
			st := status.New(codes.NotFound, "device not found")
			return nil, st.Err()
		}

		return nil, err
		
	}

	return &user.GetUserDeviceResponse{Id: device.ID ,DeviceToken: device.DeviceToken, DeviceType: device.DeviceType, UserId: device.UserID}, nil
}