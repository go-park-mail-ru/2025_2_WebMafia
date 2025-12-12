package app

import (
	"fmt"
	"os"
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
	AI                 PlaylistAI       `mapstructure:"ai"`
}

type BucketsConfig struct {
	Avatars string `mapstructure:"avatars"`
}

type PlaylistAI struct {
	AuthKey string `mapstructure:"auth_key"`
	Model   string `mapstructure:"model"`
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	configName := os.Getenv("CONFIG_FILE")
	if configName == "" {
		configName = "config.dev"
	}
	v.SetConfigName(configName)
	v.SetConfigType("yml")
	v.AddConfigPath(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("cannot read playlist config: %w", err)
	}

	if err := app.BindViperEnv(v); err != nil {
		return nil, fmt.Errorf("failed to bind env variables: %w", err)
	}

	var cfg Config

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("cannot unmarshal playlist config: %w", err)
	}

	fmt.Println("CONFIG USED =", v.ConfigFileUsed())
	fmt.Println("PLAYLIST AUTH KEY =", cfg.Playlist.AI.AuthKey)
	fmt.Println("PLAYLIST AUTH LEN =", len(cfg.Playlist.AI.AuthKey))

	return &cfg, nil
}
