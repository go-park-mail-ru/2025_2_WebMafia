package app

import (
	"fmt"
	"spotify/internal/app"
	"spotify/pkg/minio"
	"spotify/pkg/postgres"

	"github.com/spf13/viper"
)

type Config struct {
	Auth  AuthConfig      `mapstructure:"auth"`
	DB    postgres.Config `mapstructure:"db"`
	Minio minio.Config    `mapstructure:"minio"`
}

type AuthConfig struct {
	HTTP   app.HTTPConfig   `mapstructure:"http"`
	GRPC   app.GRPCConfig   `mapstructure:"grpc"`
	Logger app.LoggerConfig `mapstructure:"logger"`
}

func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath(configPath)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	app.BindViperEnv(v)

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal auth config: %w", err)
	}

	return &cfg, nil
}
