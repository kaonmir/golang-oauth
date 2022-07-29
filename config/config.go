package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var cfg *config

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Error loading config/.env file")
	}

	cfg = new(config)
	if err := env.Parse(cfg); err != nil {
		log.Printf("%+v\n", err)
	}
}

func Env() *config {
	return cfg
}
