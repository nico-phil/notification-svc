package config

import (
	"log"
	"os"
	"strconv"
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

func GetBrokerUrl() string {
	return getEnvironmentvalue("BROKER_URL")
}

func GetAppPort() int {
	portStr := getEnvironmentvalue("APP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("invalid port %s", portStr)
	}
	return port 
}

func getEnvironmentvalue(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("%s environmant variable is missing", key)
	}

	return v
}