package ws

import (
	"time"
)

type Config struct {
	WriteWait       time.Duration `mapstructure:"writeWait"`
	PongWait        time.Duration `mapstructure:"pongWait"`
	PingPeriod      time.Duration `mapstructure:"pingPeriod"`
	MaxMessageSize  int64         `mapstructure:"maxMessageSize"`
	ReadBufferSize  int           `mapstructure:"readBufferSize"`
	WriteBufferSize int           `mapstructure:"writeBufferSize"`
	SendBufferSize  int           `mapstructure:"sendBufferSize"`
}

func NewDefaultConfig() Config {
	return Config{
		WriteWait:       10 * time.Second,
		PongWait:        60 * time.Second,
		PingPeriod:      54 * time.Second,
		MaxMessageSize:  2048,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		SendBufferSize:  256,
	}
}
