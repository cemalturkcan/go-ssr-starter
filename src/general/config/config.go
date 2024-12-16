package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strconv"
)

var (
	// SERVER
	AppName = os.Getenv("APP_NAME")
	Port    = ""
	PreFork = os.Getenv("PreFork") == "true"
)

const (
	Development = "development"
	Production  = "production"
)

var (
	IsProduction  = false
	IsDevelopment = false
)

func init() {
	port, err := strconv.Atoi(os.Getenv("Port"))
	if err != nil {
		Port = ":8080"
	} else {
		Port = ":" + strconv.Itoa(port)
	}

	env := os.Getenv("ENV")
	if env == "" || env == Development {
		IsDevelopment = true
	} else {
		if env == Production {
			IsProduction = true
		}
	}
}
