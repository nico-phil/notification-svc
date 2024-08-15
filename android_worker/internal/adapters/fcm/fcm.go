package fcm

import (
	"context"
	"net/http"
	"os"

	"golang.org/x/oauth2/google"
)

type Adapter struct {
	AccessToken string
	Client *http.Client
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

// func(a *Adapter) SendRequestToFireBase(){
// 	url := ""

// 	data := map[string]any {
// 		"message": map[string]any {
// 			"token": a.AccessToken,
// 			"notification": map[string] string {
// 				"title": "hello world",
// 				"body": "This is an FCM notification message",
// 			},
// 		},
		
// 	}

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println("error marshaling data", err)
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
// 	if err!= nil {
// 		fmt.Println("error creating request", err)
// 	}

// 	req.Header.Set("Content-Type", "application/json")


// 	_, err = a.Client.Do(req)
// 	if err != nil {
// 		fmt.Println("error making the request", err)
// 	}
// 	fmt.Println("all good")
// }




