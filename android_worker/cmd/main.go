package main

import (
	"log"

	"github.com/nico-phil/notification_worker/internal/adapters/consumer"
)

func main(){

	consumerAdapter, err := consumer.NewAdapter([]string{"localhost:9092"})

	if err != nil {
		log.Fatalf("failed to connect to kafka err: %v", err)
	}

	consumerAdapter.ConsumeMessageFromQueue()

}
