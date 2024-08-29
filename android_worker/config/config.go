package config

import (
	"log"
	"os"
)

func GetBrokerUrl() string {
	return getEnvironmentvalue("BROKER_URL")
}

func GetFirebaseProjectId() string {
	return getEnvironmentvalue("FIREBASE_PROJECT_ID")
}

func getEnvironmentvalue(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("%s environmant variable is missing", key)
	}

	return v
}