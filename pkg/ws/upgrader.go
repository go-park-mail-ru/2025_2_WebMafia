package ws

import (
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

func NewUpgrader(allowedOrigins []string, config Config) websocket.Upgrader {
	return websocket.Upgrader{
		ReadBufferSize:  config.ReadBufferSize,
		WriteBufferSize: config.WriteBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			for _, allowed := range allowedOrigins {
				if strings.EqualFold(origin, allowed) {
					return true
				}
			}
			return false
		},
	}
}
