package app

import (
	"os"
	"spotify/internal/handler"
	"strings"
	"time"
)

type Config struct {
	Port            string
	ShutdownTimeout time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	CORS            handler.CORSConfig
}

func NewConfig() *Config {
	return &Config{
		Port:            getEnv("PORT", "8080"),
		ShutdownTimeout: getEnvAsDuration("SHUTDOWN_TIMEOUT", 15*time.Second),
		ReadTimeout:     getEnvAsDuration("READ_TIMEOUT", 5*time.Second),
		WriteTimeout:    getEnvAsDuration("WRITE_TIMEOUT", 10*time.Second),
		IdleTimeout:     getEnvAsDuration("IDLE_TIMEOUT", 60*time.Second),
		CORS: handler.CORSConfig{
			AllowedOrigins:   getEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"http://localhost:3030"}),
			AllowedMethods:   getEnvAsSlice("CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}),
			AllowedHeaders:   getEnvAsSlice("CORS_ALLOWED_HEADERS", []string{"Content-Type", "Authorization", "X-Requested-With"}),
			AllowCredentials: getEnvAsBool("CORS_ALLOW_CREDENTIALS", true),
		},
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

func getEnvAsSlice(key string, defaultValue []string) []string {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	return strings.Split(valueStr, ",")
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	return strings.ToLower(valueStr) == "true" || valueStr == "1"
}
