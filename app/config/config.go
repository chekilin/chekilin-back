package config

import (
	"fmt"
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Port       int    `env:"PORT" envDefault:"8080"`
	DBHost     string `env:"CHEKI_DB_HOST"`
	DBPort     int    `env:"CHEKI_DB_PORT" envDefault:"33306"`
	DBUser     string `env:"CHEKI_DB_USER"`
	DBPassword string `env:"CHEKI_DB_PASSWORD"`
	DBName     string `env:"CHEKI_DB_NAME" envDefault:"cheki"`
}

func New() (*Config, error) {
	if err := godotenv.Load("env/test.env"); err != nil {
		log.Fatal("Error loading test.env file")
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Println(err)
		return nil, err
	}

	user := os.Getenv("CHEKI_DB_USER")
	password := os.Getenv("CHEKI_DB_PASSWORD")
	cfg.DBUser = user
	cfg.DBPassword = password

	return cfg, nil
}
