package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Application struct {
	GinMode     string
	RunPort     string
	DBSource    string
	MemcachedIP string
}

func LoadConfig() Application {
	cfg := Application{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file not found")
	}

	cfg.GinMode = os.Getenv("GIN_MODE")
	cfg.RunPort = os.Getenv("RUN_PORT")
	cfg.DBSource = os.Getenv("DB_SOURCE")
	cfg.MemcachedIP = os.Getenv("MEMCACHED_IP")

	return cfg
}
