package api

import (
	"context"
	"errors"
	"testing"

	"github.com/nico-phil/notification/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockedUser struct {
	mock.Mock
}

func(u *mockedUser) Get(ctx context.Context, id int64) (domain.User, error) {
	args := u.Called(ctx, id)
	return args.Get(0).(domain.User), args.Error(1)
}

func(u *mockedUser) GetDevice(ctx context.Context, userId int64)(domain.Device, error){
	args := u.Called(ctx, userId)
	return args.Get(0).(domain.Device), args.Error(1)
}

type mockedProducer struct {
	mock.Mock
}

func(p *mockedProducer) PushMessageToQueue(topic string, message domain.PushNotification) error {
	args := p.Called(topic, message)
	return args.Error(0)
}

func(p *mockedProducer) PushMessageToQueueEmail(topic string, message domain.EmailNotification) error {
	args := p.Called(topic, message)
	return args.Error(0)
}


func(p *mockedProducer) PushMessageToQueueSMS(topic string, message domain.SMSNotification) error {
	args := p.Called(topic, message)
	return args.Error(0)
}

func TestPushNotification(t *testing.T) {
	user := new(mockedUser)
	producer := new(mockedProducer)

	user.On("GetDevice", mock.Anything, mock.Anything).Return(domain.Device{
		DeviceToken: "123",
		DeviceType: "IOS",
		UserId: 1,
	}, nil)

	producer.On("PushMessageToQueue", mock.Anything, mock.Anything).Return(nil)
	
	application := NewApplication(producer, user)
	err := application.SendPushNotification(context.Background(), domain.Notification{
		Title: "Gretting",
		Content: "Hello Dear user",
		UserId: 1,
		NotificationType: "PUSH",
	})

	assert.Nil(t, err)

}

func Test_Should_Return_Error_When_Device_Not_Found(t *testing.T){
	user := new(mockedUser)
	producer := new(mockedProducer)

	deviceCache = make(map[int64]domain.Device)
	
	user.On("GetDevice", mock.Anything, mock.Anything).Return(domain.Device{}, errors.New("device not found"))
	producer.On("PushMessageToQueue", mock.Anything, mock.Anything).Return(nil)

	application := NewApplication(producer, user)
	err := application.SendPushNotification(context.Background(), domain.Notification{
		Title: "Gretting",
		Content: "Hello Dear user",
		UserId: 1,
		NotificationType: "PUSH",
	})

	assert.EqualError(t, err, "device not found")
}

func Test_Should_Return_Error_When_PushNotificationPushToQueue_Fail(t *testing.T){
	user := new(mockedUser)
	producer := new(mockedProducer)

	user.On("GetDevice", mock.Anything, mock.Anything).Return(domain.Device{}, nil)
	producer.On("PushMessageToQueue", mock.Anything, mock.Anything).Return(errors.New("error pushing to queue"))

	application := NewApplication(producer, user)
	err := application.SendPushNotification(context.Background(), domain.Notification{
		Title: "Gretting",
		Content: "Hello Dear user",
		UserId: 1,
		NotificationType: "PUSH",
	})

	assert.EqualError(t, err, "error pushing to queue")
}

