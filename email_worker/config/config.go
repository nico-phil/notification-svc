package config

import (
	"log"
	"os"
)

func GetEnv() string {
	return getEnvironmentvalue("ENV")
}

func GetDomain() string {
	return getEnvironmentvalue("DOMAIN")
}

func GetApiToken() string {
	return getEnvironmentvalue("API_TOKEN")
}

func GetEmail() string {
	return getEnvironmentvalue("EMAIL")
}

func GetDbDSN() string{
	return getEnvironmentvalue("DSN")
}

func getEnvironmentvalue(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("%s environmant variable is missing", key)
	}

	return v
}