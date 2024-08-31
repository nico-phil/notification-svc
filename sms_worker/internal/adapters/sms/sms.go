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

func(a *Adapter) SendSMSNotification(body, from, to string) error{ 
	sms_url := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", config.GetTwiolioAccountSID())

	if config.GetEnv() == "development" {
		from = config.GetFrom()
		to = config.GetTO()
	}

	form := url.Values{}
	form.Add("Body", body)
	form.Add("From", from)
	form.Add("To", to) 

	bodyReader := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", sms_url, bodyReader)
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