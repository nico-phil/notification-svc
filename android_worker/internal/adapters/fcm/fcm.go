package fcm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
)

type Adapter struct {
	AccessToken string
	Client http.Client
}

func(a *Adapter) GenerateToken() error {
	
	serviceAccountFile := "./notification-6b719-firebase-adminsdk-r6d81-e1e692321c.json"

	data, err := os.ReadFile(serviceAccountFile)
	if err != nil {
		return err
	}

	creds, err := google.CredentialsFromJSON(context.Background(), data, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return err
	}
	token, err := creds.TokenSource.Token()
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

func(a *Adapter) SendRequestToFireBase(title, body, token string){
	url := fmt.Sprintf("https://fcm.googleapis.com/v1/projects/%s/messages:send", os.Getenv("FIREBASE_PROJECT_ID"))

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
		fmt.Println("error marshaling data", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err!= nil {
		fmt.Println("error creating request", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken) )


	_, err = a.Client.Do(req)
	if err != nil {
		fmt.Println("error making the request", err)
	}
	fmt.Println("all good")
}




