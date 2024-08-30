package sms

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/nico-phi/notification/sms_worker/config"
)

type Adapter struct {
	Client *http.Client
}

type Message struct {
	Body string
	From string
	To string
}

func(a *Adapter) SendSMSNotification(message Message) error{
	sms_url := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", config.GetTwiolioAccountSID())

	if config.GetEnv() == "development" {
		message.From = config.GetFrom()
		message.To = config.GetTO()
	}

	form := url.Values{}
	form.Add("Body", message.Body)
	form.Add("From", message.From)
	form.Add("To", message.To) 

	body := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", sms_url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",config.GetTwiolioAccountSID(),config.GetTwilioAuthtoken()))))

	resp, err := a.Client.Do(req)

	fmt.Println("StatusCode=",resp.StatusCode)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return errors.New("bad request")
	}

	return nil
}