package app

import (
	"fmt"
	"spotify/internal/app"
	"spotify/pkg/postgres"

	"github.com/spf13/viper"
)

type Config struct {
	Catalog CatalogConfig   `mapstructure:"catalog"`
	DB      postgres.Config `mapstructure:"db"`
}

type CatalogConfig struct {
	HTTP   app.HTTPConfig   `mapstructure:"http"`
	GRPC   GRPCConfig       `mapstructure:"grpc"`
	Logger app.LoggerConfig `mapstructure:"logger"`
}

type GRPCConfig struct {
	app.GRPCConfig `mapstructure:",squash"`
	Clients        ClientsConfig `mapstructure:"clients"`
}

type ClientsConfig struct {
	Auth string `mapstructure:"auth"`
}

func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath(configPath)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := app.BindViperEnv(v); err != nil {
		return nil, fmt.Errorf("failed to bind env variables: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal catalog config: %w", err)
	}

	return &cfg, nil
}
