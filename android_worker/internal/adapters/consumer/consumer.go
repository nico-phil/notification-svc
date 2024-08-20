package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"github.com/nico-phil/notification_worker/internal/ports"
)

type Adapter struct {
	consumer sarama.Consumer
	Topic string
	FCM ports.FCMPort

}

type Notification struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

type Device struct{
	ID int	`json:"id"`
	DeviceToken string `json:"device_token"`
	DeviceType string 	`json:"device_type"`
}

type PushNotification struct {
	Notification Notification 	`json:"notification"`
	Device Device `json:"device"`
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

func(a Adapter) ConsumeMessageFromQueue(){
	partitionConsumer, err :=  a.consumer.ConsumePartition(a.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Println("failed to create partition consumer")
	}

	defer func(){
		if err := partitionConsumer.Close(); err!= nil {
			log.Println("failed to close consumer")
		}
	}()

	fmt.Println("start consumming message")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	doneCh := make(chan struct{})
	
	msgCnt := 0

	go func() {
		for {
			select {
			case err := <-partitionConsumer.Errors():
				fmt.Println(err)
			case msg := <-partitionConsumer.Messages():
				msgCnt++
				fmt.Printf("Received Notification Count %d: | Topic(%s) \n", msgCnt, string(msg.Topic))
				a.ProcessMessage(msg)
				
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	ts, _ := a.consumer.Topics()
	fmt.Println("total topics:", ts)

}


func(a *Adapter) ProcessMessage(msg *sarama.ConsumerMessage){
	value := msg.Value
 
	var pushNotification PushNotification
	err := json.Unmarshal(value, &pushNotification)
	if err != nil {
		fmt.Println("failed unmarshaling data", err)
	}
			
	count := 3
	for count > 0 {
		err  = a.FCM.SendNotification(pushNotification.Notification.Title, pushNotification.Notification.Content, pushNotification.Device.DeviceToken)
		if err != nil {
			count--
			log.Println(err)
			time.Sleep(2 * time.Second)
		}else {
			fmt.Println("suceess process notification", pushNotification.Notification.Content)
			break;
		}
	} 
	
}


