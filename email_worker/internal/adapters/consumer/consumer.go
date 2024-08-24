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
	"github.com/nico-phil/email_worker/internal/ports"
)

type Adapter struct {
	consumer sarama.Consumer
	Topic string
	Mail ports.MailPort

}

type EmailNotification struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Email string `json:"email"`
}

func NewAdapter(mailPort ports.MailPort, brokers []string) (*Adapter, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}
	
	return &Adapter{consumer: consumer, Topic: "EMAIL_QUEUE", Mail: mailPort }, nil
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
				fmt.Printf("Received message Count %d: | Topic(%s) \n", msgCnt, string(msg.Topic))
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
 
	var payload EmailNotification
	err := json.Unmarshal(value, &payload)
	if err != nil {
		fmt.Println("failed to unmarshal data", err)
	}
			
	count := 3
	for count > 0 {
		err  := a.Mail.SendRequestToMailSender(payload.Title, payload.Content, payload.Email)
		if err != nil {
			count--
			log.Println(err)
			time.Sleep(2 * time.Second)
		}else {
			fmt.Println("suceess process notification", )
			break;
		}
	} 
	
}


