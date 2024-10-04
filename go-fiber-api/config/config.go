package config

import (
	"log"
	"os"
)

var AppConfig struct {
	Port string
}

func LoadConfig() {
	AppConfig.Port = os.Getenv("PORT")
	// if AppConfig.Port == "3000" {
	// 	log.Fatal("PORT is not set in environment variables")
	// } else
	if AppConfig.Port == "" {
		log.Fatal("PORT is not set in environment variables")
	}
}
