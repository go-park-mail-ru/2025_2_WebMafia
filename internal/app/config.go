package app

import (
	"spotify/internal/middleware"
	"time"
)

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
