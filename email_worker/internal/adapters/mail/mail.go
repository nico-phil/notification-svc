package mail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nico-phil/email_worker/config"
)

type Mail struct {
	API_TOKEN string
	Client http.Client
}

type Payload struct {
	From FromTo `json:"from"`
	To FromTo `json:"to"`
	Subject string `json:"subject"`
	Text string `json:"text"`
	Html string `json:"html"`
}

type FromTo struct {
	Email string `json:"email"`
}


func(m *Mail) SendRequestToMailSender(subject, text, email string)error {
	url := "https://api.mailersend.com/v1/email"

	payload := Payload {
		From: FromTo{Email: config.GetDomain()},
		To: FromTo{Email: config.GetEmail()},
		Subject: subject,
		Text: text,
		Html: text,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+ m.API_TOKEN)
	req.Header.Set("Content-Type", "application/json")

	res, err := m.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode !=  http.StatusOK {
		fmt.Println(res.Status)
	}

	return nil
}