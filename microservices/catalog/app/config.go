package app

import (
	"fmt"
	"spotify/internal/app"
	"spotify/pkg/postgres"

	"github.com/spf13/viper"
)

type Config struct {
	Catalog CatalogConfig `mapstructure:"catalog"`
	DB      postgres.Config
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

	var raw struct {
		Catalog CatalogConfig              `mapstructure:"catalog"`
		DB      map[string]postgres.Config `mapstructure:"db"`
	}

	if err := v.Unmarshal(&raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal catalog config: %w", err)
	}

	cfg := &Config{
		Catalog: raw.Catalog,
		DB:      raw.DB["catalog"],
	}

	return cfg, nil
}
