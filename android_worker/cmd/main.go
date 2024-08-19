package main

import (
	"log"

	"github.com/nico-phil/notification_worker/internal/adapters/consumer"
	"github.com/nico-phil/notification_worker/internal/adapters/fcm"
)

func main(){

	fcmAdapter := &fcm.Adapter{}

	err := fcmAdapter.GenerateToken()
	if err != nil {
		log.Printf("failed generating google cloud patform access token : %v", err)
	}

	consumerAdapter, err := consumer.NewAdapter(fcmAdapter, []string{"localhost:9092"})

	if err != nil {
		log.Fatalf("failed to connect to kafka broker err: %v", err)
	}

	consumerAdapter.ConsumeMessageFromQueue()

}

