package main

import (
	"log"
	"net/http"

	"github.com/nico-phi/notification/sms_worker/config"
	"github.com/nico-phi/notification/sms_worker/internal/adapters/consumer"
	"github.com/nico-phi/notification/sms_worker/internal/adapters/sms"
)

func main(){
	smsAdapter := &sms.Adapter{Client: &http.Client{}}
	 consumerAdapter, err := consumer.NewAdapter(smsAdapter, []string{config.GetBrokerUrl()})
	 if err != nil {
		log.Fatal("failed to connect to borker", err)
	 }
	
	 consumerAdapter.ConsumeMessageFromQueue()
}