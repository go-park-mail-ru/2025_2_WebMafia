package app

import (
	"fmt"
	"spotify/internal/app"
	"spotify/pkg/minio"
	"spotify/pkg/postgres"

	"github.com/spf13/viper"
)

type Config struct {
	Playlist PlaylistConfig  `mapstructure:"playlist"`
	DB       postgres.Config `mapstructure:"db"`
	Minio    minio.Config    `mapstructure:"minio"`
}

type ClientsConfig struct {
	Catalog string `mapstructure:"catalog"`
	Auth    string `mapstructure:"auth"`
}

type PlaylistConfig struct {
	HTTP               app.HTTPConfig   `mapstructure:"http"`
	Logger             app.LoggerConfig `mapstructure:"logger"`
	AllowedAvatarTypes []string         `mapstructure:"allowed_avatar_types"`
	Buckets            BucketsConfig    `mapstructure:"buckets"`
	Clients            ClientsConfig    `mapstructure:"clients"`
}

type BucketsConfig struct {
	Avatars string `mapstructure:"avatars"`
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("cannot read playlist config: %w", err)
	}

	app.BindViperEnv(v)

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("cannot unmarshal playlist config: %w", err)
	}

	return &cfg, nil
}
