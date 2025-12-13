package ws

import (
	"context"
	"spotify/internal/middleware"
	"spotify/pkg/logger"
	"time"

	"github.com/gorilla/websocket"
)

type Handler interface {
	HandleMessage(ctx context.Context, client *Client, message []byte)
}

type Client struct {
	topic   string
	id      string
	hub     *Hub
	conn    *websocket.Conn
	send    chan []byte
	config  Config
	logger  logger.Logger
	handler Handler
}

func NewClient(
	topic string,
	id string,
	hub *Hub,
	conn *websocket.Conn,
	logger logger.Logger,
	handler Handler,
	config Config,
) *Client {
	return &Client{
		topic:   topic,
		id:      id,
		hub:     hub,
		conn:    conn,
		send:    make(chan []byte, config.SendBufferSize),
		logger:  logger,
		handler: handler,
		config:  config,
	}
}

func (c *Client) ID() string {
	return c.id
}

func (c *Client) Topic() string {
	return c.topic
}

func (c *Client) SendMessage(data []byte) {
	select {
	case c.send <- data:
	default:
		c.logger.Warnf("Client %s send channel full, dropping message", c.id)
	}
}

func (c *Client) ReadPump() {
	const op = "pkg.ws.Client.ReadPump"

	defer func() {
		c.hub.Unregister(c)
		c.conn.Close()
	}()

	c.conn.SetReadLimit(c.config.MaxMessageSize)

	if err := c.conn.SetReadDeadline(time.Now().Add(c.config.PongWait)); err != nil {
		c.logger.Errorf("[%s]: client %s: set read deadline failed: %v", op, c.id, err)
		return
	}

	c.conn.SetPongHandler(func(string) error {
		return c.conn.SetReadDeadline(time.Now().Add(c.config.PongWait))
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.logger.Warnf("[%s]: client %s: unexpected close error: %v", op, c.id, err)
			} else {
				c.logger.Debugf("[%s]: client %s: connection closed", op, c.id)
			}
			break
		}

		ctx := middleware.ContextWithLogger(context.Background(), c.logger)
		c.handler.HandleMessage(ctx, c, message)
	}
}

func (c *Client) WritePump() {
	const op = "pkg.ws.Client.WritePump"

	ticker := time.NewTicker(c.config.PingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if err := c.conn.SetWriteDeadline(time.Now().Add(c.config.WriteWait)); err != nil {
				c.logger.Errorf("[%s]: client %s: set write deadline failed: %v", op, c.id, err)
				return
			}

			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					c.logger.Warnf("[%s]: client %s: failed to write close message: %v", op, c.id, err)
				}
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				c.logger.Errorf("[%s]: client %s: next writer failed: %v", op, c.id, err)
				return
			}

			if _, err := w.Write(message); err != nil {
				c.logger.Errorf("[%s]: client %s: write message failed: %v", op, c.id, err)
				return
			}

			n := len(c.send)
			for i := 0; i < n; i++ {
				if _, err := w.Write([]byte{'\n'}); err != nil {
					c.logger.Errorf("[%s]: client %s: write delimiter failed: %v", op, c.id, err)
					return
				}
				if _, err := w.Write(<-c.send); err != nil {
					c.logger.Errorf("[%s]: client %s: write buffered message failed: %v", op, c.id, err)
					return
				}
			}

			if err := w.Close(); err != nil {
				c.logger.Errorf("[%s]: client %s: writer close failed: %v", op, c.id, err)
				return
			}

		case <-ticker.C:
			if err := c.conn.SetWriteDeadline(time.Now().Add(c.config.WriteWait)); err != nil {
				c.logger.Errorf("[%s]: client %s: set ping write deadline failed: %v", op, c.id, err)
				return
			}
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.logger.Warnf("[%s]: client %s: write ping failed: %v", op, c.id, err)
				return
			}
		}
	}
}
