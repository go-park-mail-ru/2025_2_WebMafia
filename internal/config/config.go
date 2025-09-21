package config
import (
	"os"
	"time"
)

type Config struct {
	Port            string
	ShutdownTimeout time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
}

func New() *Config {
	return &Config{
		Port:            getEnv("PORT", "8080"),
		ShutdownTimeout: getEnvAsDuration("SHUTDOWN_TIMEOUT", 15*time.Second),
		ReadTimeout:     getEnvAsDuration("READ_TIMEOUT", 5*time.Second),
		WriteTimeout:    getEnvAsDuration("WRITE_TIMEOUT", 10*time.Second),
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