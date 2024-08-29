package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string{
	return getEnvironmentValue("ENV")
}

func GetDSN() string {
	return getEnvironmentValue("DSN")
}

func GetAppPort() int {
	portStr := getEnvironmentValue("APP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("invalid port %s", portStr)
	}
	return port 
}


func getEnvironmentValue(key string) string {
	v := os.Getenv(key)
	if v == ""{
		log.Fatal("error missing env variable", key)
	}

	return v
}