package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/nico-phil/notification_worker/config"
	"github.com/nico-phil/notification_worker/internal/adapters/consumer"
	"github.com/nico-phil/notification_worker/internal/adapters/fcm"
)

func main(){

	tr := &http.Transport {TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	fcmAdapter := &fcm.Adapter{ Client: &http.Client{Transport: tr}}

	err := fcmAdapter.GenerateToken()
	if err != nil {
		log.Fatalf("failed to generate google cloud patform access token : %v", err)
	}

	consumerAdapter, err := consumer.NewAdapter(fcmAdapter, []string{config.GetBrokerUrl()})

	if err != nil {
		log.Fatalf("failed to connect to kafka broker err: %v", err)
	}

	consumerAdapter.ConsumeMessageFromQueue()

}

