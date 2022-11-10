package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"todos-api-go/common/parser"
)

type Settings struct {
	GO_ENV string
	PORT   string

	// Database
	DB_HOST     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
}

func NewSetting() Settings {
	var err error
	var GO_ENV string
	if GO_ENV = os.Getenv("GO_ENV"); GO_ENV == "local" {
		err = godotenv.Load(".env.dev.local")
	} else {
		err = godotenv.Load(".env")
	}
	if err != nil {
		log.Fatal(err)
	}
	return Settings{
		GO_ENV: GO_ENV,
		PORT:   parser.GetEnv("PORT", ""),

		DB_HOST:     parser.GetEnv("DB_HOST", ""),
		DB_USERNAME: parser.GetEnv("DB_USERNAME", ""),
		DB_PASSWORD: parser.GetEnv("DB_PASSWORD", ""),
		DB_DATABASE: parser.GetEnv("DB_DATABASE", ""),
	}
}
