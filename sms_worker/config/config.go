package config

import (
	"log"
	"os"
)

func GetEnv() string{
	return getEnvironmentValue("ENV")
}

func GetFrom()string {
	return getEnvironmentValue("FROM")
}

func GetTO()string {
	return getEnvironmentValue("TO")
}

func GetTwiolioAccountSID()string {
	return getEnvironmentValue("TWILIO_ACCOUNT_SID")
}

func GetTwilioAuthtoken() string {
	return getEnvironmentValue("TWILIO_AUTH_TOKEN")
}



func getEnvironmentValue(key string) string {
	v := os.Getenv(key)
	if v == ""{
		log.Fatal("missing environment variable", key)
	}

	return v
}
