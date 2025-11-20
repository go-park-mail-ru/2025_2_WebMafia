package app

import (
	"fmt"
	"spotify/internal/middleware"
	"spotify/pkg/minio"
	"spotify/pkg/postgres"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Catalog ServiceConfig   `mapstructure:"catalog"`
	Auth    ServiceConfig   `mapstructure:"auth"`
	DB      postgres.Config `mapstructure:"db"`
	Minio   minio.Config    `mapstructure:"minio"`
}

type ServiceConfig struct {
	HTTP    HTTPConfig    `mapstructure:"http"`
	GRPC    GRPCConfig    `mapstructure:"grpc"`
	Logger  LoggerConfig  `mapstructure:"logger"`
	Clients ClientsConfig `mapstructure:"clients"`
}

type ClientsConfig struct {
	Auth string `mapstructure:"auth"`
}

type HTTPConfig struct {
	Port               string                `mapstructure:"port"`
	ReadTimeout        time.Duration         `mapstructure:"readTimeout"`
	WriteTimeout       time.Duration         `mapstructure:"writeTimeout"`
	IdleTimeout        time.Duration         `mapstructure:"idleTimeout"`
	ShutdownTimeout    time.Duration         `mapstructure:"shutdownTimeout"`
	AllowedAvatarTypes []string              `mapstructure:"allowedAvatarTypes"`
	Auth               AuthConfig            `mapstructure:"auth"`
	CORS               middleware.CORSConfig `mapstructure:"cors"`
}

type GRPCConfig struct {
	Port string `mapstructure:"port"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
	Mode  string `mapstructure:"mode"`
}

type AuthConfig struct {
	JWT  JWTConfig  `mapstructure:"jwt"`
	CSRF CSRFConfig `mapstructure:"csrf"`
}

type JWTConfig struct {
	SecretKey      string        `mapstructure:"secretKey"`
	AccessTokenTTL time.Duration `mapstructure:"accessTokenTTL"`
}

type CSRFConfig struct {
	SecretKey string        `mapstructure:"secretKey"`
	TokenTTL  time.Duration `mapstructure:"tokenTTL"`
}

func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath(configPath)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
