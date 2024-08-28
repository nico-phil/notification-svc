package api

import (
	"context"
	"errors"
	"testing"

	"github.com/nico-phil/notification/user-svc/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockedDB struct {
	mock.Mock
}

func(d *mockedDB) SaveUser(ctx context.Context, user *domain.User) error{
	args := d.Called(ctx, user)
	return args.Error(0)
}

func(d *mockedDB) GetUser(ctx context.Context, id int64) (domain.User, error){
	args := d.Called(ctx, id)
	return args.Get(0).(domain.User), args.Error(1)
}

func(d *mockedDB) SaveDevice(ctx context.Context, device *domain.Device) error {
	args := d.Called(ctx, device)
	return args.Error(0)
}

func(d * mockedDB) GetUserDevice(ctx context.Context, userId int64)(domain.Device, error){
	args := d.Called(ctx, userId)
	return args.Get(0).(domain.Device), args.Error(1)
}

func Test_Create(t * testing.T){
	db := new(mockedDB)
	db.On("SaveUser", mock.Anything, mock.Anything).Return(nil)

	u := domain.User{
		Firstname: "jhon",
		Lastname: "doe",
		Email: "jhondoe@gmai.com",
		Password: "123456",
	}
	u.EncriptPassword()
	application := NewApplication(db)
	err := application.Create(context.Background(), &u)

	assert.Nil(t,err)

}

func Test_Create_Should_Retur_Error_when_db_Persistent_Fail(t *testing.T){
	db := new(mockedDB)
	db.On("SaveUser", mock.Anything, mock.Anything).Return(errors.New("failed to save data"))

	u := domain.User{
		Firstname: "jhon",
		Lastname: "doe",
		Email: "jhondoe@gmai.com",
		Password: "123456",
	}
	u.EncriptPassword()
	application := NewApplication(db)
	err := application.Create(context.Background(), &u)

	assert.EqualError(t, err, "failed to save data")
}

func Test_Get(t *testing.T){
	db := new(mockedDB)
	db.On("GetUser", mock.Anything, mock.Anything).Return(domain.User{}, nil)

	application := NewApplication(db)
	u, err := application.db.GetUser(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, u, domain.User{})
}

func Test_Get_Should_Return_Error_When_User_Not_Found(t *testing.T) {
	db := new(mockedDB)
	db.On("GetUser", mock.Anything, mock.Anything).Return(domain.User{}, errors.New("user not found"))

	application := NewApplication(db)
	_, err := application.db.GetUser(context.Background(), 1)

	assert.EqualError(t, err, "user not found")
}


func Test_CreateDevice(t *testing.T){
	db := new(mockedDB)
	db.On("SaveDevice", mock.Anything, mock.Anything).Return(nil)

	application := NewApplication(db)
	err := application.CreateDevice(context.Background(), &domain.Device{
		DeviceToken: "1213",
		DeviceType: "IOS",
		UserID: 1,
	})

	assert.Nil(t, err)
}

func Test_CreateDevice_Should_Return_Error_when_Db_Persistent_Fail(t *testing.T){
	db := new(mockedDB)
	db.On("SaveDevice", mock.Anything, mock.Anything).Return(errors.New("failed to save device"))

	application := NewApplication(db)
	err := application.CreateDevice(context.Background(), &domain.Device{
		DeviceToken: "1213",
		DeviceType: "IOS",
		UserID: 1,
	})

	assert.EqualError(t, err, "failed to save device")
}

func Test_GetUserDevice(t *testing.T){
	db := new(mockedDB)
	db.On("GetUserDevice", mock.Anything, mock.Anything).Return(domain.Device{}, nil)
	
	application := NewApplication(db)
	d, err := application.GetUserDevice(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, d, domain.Device{})

}

func Test_GetUserDevice_Should_Return_Error_When_Device_Not_Found(t *testing.T){
	db := new(mockedDB)
	db.On("GetUserDevice", mock.Anything, mock.Anything).Return(domain.Device{}, errors.New("device not found"))
	
	application := NewApplication(db)
	_, err := application.GetUserDevice(context.Background(), 1)

	assert.EqualError(t, err, "device not found")
}



