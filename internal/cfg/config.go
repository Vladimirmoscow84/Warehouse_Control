package cfg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURI   string
	ServerAddress string
	JWTSecret     string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		DatabaseURI:   os.Getenv("DATABASE_URI"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
	}
	if cfg.DatabaseURI == "" {
		log.Fatal("DATABASE_URI is not set")
	}
	if cfg.ServerAddress == "" {
		log.Fatal("SERVER_ADDRESS is not set")
	}
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	return cfg
}
