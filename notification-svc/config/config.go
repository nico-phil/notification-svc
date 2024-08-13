package config

import (
	"log"
	"os"
)

func GetEnv() string {
	return getEnvironmentvalue("ENV")
}

func getEnvironmentvalue(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("%s environmant variable is missing", key)
	}

	return v
}