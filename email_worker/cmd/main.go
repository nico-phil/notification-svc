package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/nico-phil/email_worker/config"
	"github.com/nico-phil/email_worker/internal/adapters/consumer"
	"github.com/nico-phil/email_worker/internal/adapters/mail"
)



func main() {
	tr := &http.Transport {TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	mail := &mail.Mail{
		API_TOKEN: config.GetApiToken(),
		Client: &http.Client{Transport: tr},
	}
	consumerAdapter, err := consumer.NewAdapter(mail, []string{config.GetBrokerUrl()})

	if err != nil {
		log.Fatalf("failed to connect to kafka broker err: %v", err)
	}

	consumerAdapter.ConsumeMessageFromQueue()
}