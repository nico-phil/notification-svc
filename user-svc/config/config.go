package config

import (
	"log"
	"os"
)

func GetEnv() string{
	return getEnvironmentVariable("ENV")
}


func getEnvironmentVariable(key string) string {
	v := os.Getenv(key)
	if v == ""{
		log.Fatal("error missing env variable", key)
	}

	return v
}