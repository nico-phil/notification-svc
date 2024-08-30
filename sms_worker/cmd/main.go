package main

import (
	"net/http"

	"github.com/nico-phi/notification/sms_worker/internal/adapters/sms"
)

func main(){
	smsAdapter := sms.Adapter{Client: &http.Client{}}
	smsAdapter.SendSMSNotification(sms.Message{Body: "TEST MESSAGE FROM MAIN"})
}