package consumer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/nico-phil/notification_worker/internal/ports"
)

type Adapter struct {
	consumer sarama.Consumer
	Topic string
	FCM ports.FCMPort

}

type Notification struct {
	Content string `json:"content"`
	From string `json:"from"`
	To string `json:"to"`
	NotifType string `json:"notif_type"`
	Device Device `json:"device"`

}

type Device struct{
	DeviceToken string `json:"device_token"`
	DeviceType string 	`json:"device_type"`
}

func NewAdapter(fcmPort ports.FCMPort, brokers []string) (*Adapter, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}
	
	return &Adapter{consumer: consumer, Topic: "ANDROID_QUEUE", FCM: fcmPort }, nil
}

func(a Adapter) ConsumeMessageFromQueue() error{
	partitionConsumer, err :=  a.consumer.ConsumePartition(a.Topic, 0, sarama.OffsetOldest)
	
	if err != nil {
		return err
	}

	defer partitionConsumer.Close()

	fmt.Println("start consumming message")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})
	
	msgCnt := 0

	go func() {
		for {
			select {
			case err := <-partitionConsumer.Errors():
				fmt.Println(err)
			case msg := <-partitionConsumer.Messages():
				msgCnt++
				value := msg.Value
				var notification Notification
				err := json.Unmarshal(value, &notification)
				if err != nil {
					fmt.Println("failed unmarshaling data", err)
				}
				a.FCM.SendNotification(notification.From, notification.Content, notification.Device.DeviceToken)
				fmt.Printf("Received Notification Count %d: | Topic(%s) | Message(%s) \n", msgCnt, string(msg.Topic), notification.Content)
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	ts, _ := a.consumer.Topics()
	fmt.Println("total topics:", ts)


	return nil
}







