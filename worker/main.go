package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main(){

	topic := "IOS_QUEUE"
	msgCnt := 0

	worker, err := ConnectConsumer([]string{"localhost:9092"})
	if err != nil {
		log.Fatalf("failed to connect consumer %v", err)
	}

	defer func() {
		if err := worker.Close(); err != nil {
			log.Fatalln("opps somthing went wrong", err)
		}
	}()

	consumerPartition, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	defer consumerPartition.Close()

	fmt.Println("consumer started")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumerPartition.Errors():
				fmt.Println(err)
			case msg := <-consumerPartition.Messages():
				msgCnt++
				fmt.Printf("Received Notification Count %d: | Topic(%s) | Message(%s) \n", msgCnt, string(msg.Topic), string(msg.Value))
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCnt, "messages")
	ts, _ := worker.Topics()
	fmt.Println("total topics:", ts)
}

func ConnectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	return sarama.NewConsumer(brokers, config)
}