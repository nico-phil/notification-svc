package producer

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/nico-phil/notification/internal/application/core/domain"
)


type Adapter struct {
	producer sarama.SyncProducer
}

func NewAdapter(brokers []string) (*Adapter, error){

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return &Adapter{producer: producer}, nil
}

func (a *Adapter) PushMessageToQueue(topic string, message domain.PushNotification) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: message,
	}

	fmt.Println(message)

	// defer func(){
	// 	err :=  a.producer.Close()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	partition, offset, err:= a.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Notification is stored in topic (%s)/partition/(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}

