package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       int    `env:"PORT" envDefault:"80"`
	DBHost     string `env:"CHEKI_DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"CHEKI_DB_PORT" envDefault:"33306"`
	DBUser     string `env:"CHEKI_DB_USER" envDefault:"cheki"`
	DBPassword string `env:"CHEKI_DB_PASSWORD" envDefault:"cheki"`
	DBName     string `env:"CHEKI_DB_NAME" envDefault:"cheki"`
}

func New() (*Config, error) {
	err := godotenv.Load("./env/.env.dev")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{}
	portStr := os.Getenv("PORT")
	fmt.Println(portStr)
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Error converting PORT to int: %v", err)
	}
	cfg.Port = port

	return cfg, nil
}
