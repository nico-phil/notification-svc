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


func(m *Mail) SendRequestToMailSender()error {
	url := "https://api.mailersend.com/v1/email"

	payload := Payload {
		From: FromTo{Email: config.GetDomain()},
		To: FromTo{Email: "nphilibert17@gmail.com"},
		Subject: "Hello from Nico",
		Text: "Greetings from the team, you got this message through MailerSend",
		Html: "Greetings from the team, you got this message through MailerSend",
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

	fmt.Println("req status:",res.Status)

	return nil
}