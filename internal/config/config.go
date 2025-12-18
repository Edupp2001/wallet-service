package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() *Config {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("cannot load config.env")
	}

	return &Config{
		Port: os.Getenv("APP_PORT"),
	}
}
