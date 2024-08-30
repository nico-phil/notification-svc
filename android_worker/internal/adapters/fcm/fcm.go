package fcm

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/nico-phil/notification_worker/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Adapter struct {
	AccessToken string
	Client *http.Client
}

func(a *Adapter) GenerateToken() error {
	
	serviceAccountFile := "/usr/src/app/notification-6b719-firebase-adminsdk-r6d81-e1e692321c.json"

	data, err := os.ReadFile(serviceAccountFile)
	if err != nil {
		return err
	}

	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	customClient := &http.Client{Transport: transport}


	// creds, err := google.CredentialsFromJSON(context.Background(), data, "https://www.googleapis.com/auth/cloud-platform", option.WithHTTPClient(httpClient))
	// if err != nil {
	// 	return err
	// }

	config, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return err
	}

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, customClient)
	tokenSource := config.TokenSource(ctx)

	token, err := tokenSource.Token()
	if err != nil {
		return err
	}

	a.AccessToken = token.AccessToken

	return nil
}

type MessagePayload struct {
	Message Message `json:"message"`
}

type Message struct {
	Token string `json:"token"`
	Notification Notification `json:"notification"`
}

type Notification struct {
	Title string `json:"title"`
	Body string `json:"body"`
}

func(a *Adapter) SendNotification(title, body, token string) error{
	url := fmt.Sprintf("https://fcm.googleapis.com/v1/projects/%s/messages:send", config.GetFirebaseProjectId())

	data := MessagePayload {
		Message: Message{
			Token: token,
			Notification: Notification{
				Title: title,
				Body:  body,
			},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err!= nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken) )


	_, err = a.Client.Do(req)
	if err != nil {
		return err
	}
	return nil
}




