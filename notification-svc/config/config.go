package config

import (
	"log"
	"os"
)

func GetEnv() string {
	return getEnvironmentvalue("ENV")
}

func GetDbDSN() string{
	return getEnvironmentvalue("DSN")
}

func GetUserServiceUrl() string{
	return getEnvironmentvalue("USER_SERVICE_URL")
}

func getEnvironmentvalue(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("%s environmant variable is missing", key)
	}

	return v
}