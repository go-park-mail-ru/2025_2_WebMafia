package app

import (
	"log"
	"os"

	"spotify/internal/middleware"

	"spotify/pkg/postgres"

	"strconv"
	"strings"
	"time"

	"spotify/pkg/logger"

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
	CORS            middleware.CORSConfig
	DB              postgres.Config
	Logger          logger.Config
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
		CORS: middleware.CORSConfig{
			AllowedOrigins:   getEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"http://localhost:8090"}),
			AllowedMethods:   getEnvAsSlice("CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}),
			AllowedHeaders:   getEnvAsSlice("CORS_ALLOWED_HEADERS", []string{"Content-Type", "Authorization", "X-Requested-With"}),
			AllowCredentials: getEnvAsBool("CORS_ALLOW_CREDENTIALS", true),
		},
		DB: postgres.Config{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnv("DB_PORT", "5432"),
			User:            getEnv("DB_USER", "myuser"),
			Password:        getEnv("DB_PASSWORD", "mypassword"),
			DBName:          getEnv("DB_NAME", "mydb"),
			MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 25),
			ConnMaxLifetime: getEnvAsDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
		},
		Logger: logger.Config{
			Level: getEnv("LOGGER_LEVEL", logger.LevelInfo),
			Mode:  getEnv("LOGGER_MODE", logger.ModeDev),
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

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
