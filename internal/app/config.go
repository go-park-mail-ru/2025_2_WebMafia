package app

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	ShutdownTimeout time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	AccessTokenTTL  time.Duration
	JWTSecretKey    string
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("no .env file found, trying .env.dev")
		if err := godotenv.Load(".env.dev"); err != nil {
			log.Println("no .env or .env.dev file found")
		}
	}
	return &Config{
		Port:            getEnv("PORT", "8080"),
		ShutdownTimeout: getEnvAsDuration("SHUTDOWN_TIMEOUT", 15*time.Second),
		ReadTimeout:     getEnvAsDuration("READ_TIMEOUT", 5*time.Second),
		WriteTimeout:    getEnvAsDuration("WRITE_TIMEOUT", 10*time.Second),
		IdleTimeout:     getEnvAsDuration("IDLE_TIMEOUT", 60*time.Second),
		AccessTokenTTL:  getEnvAsDuration("ACCESS_TOKEN_TTL", 720*time.Hour),
		JWTSecretKey:    getEnv("JWT_SECRET_KEY", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	if duration, err := time.ParseDuration(valueStr); err == nil {
		return duration
	}
	return defaultValue
}
