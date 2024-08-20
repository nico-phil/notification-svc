package main

import (
	"log"

	"github.com/nico-phil/email_worker/internal/adapters/consumer"
	"github.com/nico-phil/email_worker/internal/adapters/mail"
)



func main() {
	mail := &mail.Mail{
		API_TOKEN: "",
	}
	consumerAdapter, err := consumer.NewAdapter(mail, []string{"localhost:9092"})

	if err != nil {
		log.Fatalf("failed to connect to kafka broker err: %v", err)
	}

	consumerAdapter.ConsumeMessageFromQueue()
}