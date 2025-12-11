package http

import (
	"context"
	"spotify/microservices/catalog/dto"
	"spotify/pkg/logger"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/mailru/easyjson"
)

const (
	writeWait          = 10 * time.Second
	pongWait           = 60 * time.Second
	pingPeriod         = (pongWait * 9) / 10
	maxMessageSize     = 2048
	maxTextLength      = 1000
	saveCommentTimeout = 5 * time.Second
)

type Client struct {
	hub     *Hub
	conn    *websocket.Conn
	send    chan []byte
	service IService
	logger  logger.Logger
	trackID uuid.UUID
	userID  uuid.UUID
}

func (c *Client) readPump() {
	const op = "delivery.http.ws_client.readPump"

	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		c.logger.Errorf("[%s]: failed to set read deadline: %v", op, err)
		return
	}
	c.conn.SetPongHandler(func(string) error {
		return c.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.logger.Errorf("[%s]: websocket error: %v", op, err)
			}
			break
		}

		var req dto.PostCommentRequest
		if err := easyjson.Unmarshal(message, &req); err != nil {
			c.logger.Warnf("[%s]: invalid json format: %v", op, err)
			continue
		}

		if len(req.Text) == 0 || len(req.Text) > maxTextLength {
			c.logger.Warnf("[%s]: invalid text length: %d", op, len(req.Text))
			continue
		}

		req.TrackID = c.trackID.String()

		ctx, cancel := context.WithTimeout(context.Background(), saveCommentTimeout)
		createdComment, err := c.service.PostComment(ctx, c.userID, req)
		cancel()

		if err != nil {
			c.logger.Errorf("[%s]: failed to save comment: %v", op, err)
			continue
		}

		responseBytes, err := easyjson.Marshal(createdComment)
		if err != nil {
			c.logger.Errorf("[%s]: failed to marshal response: %v", op, err)
			continue
		}

		c.hub.broadcast <- Message{
			TrackID: c.trackID,
			Data:    responseBytes,
		}
	}
}

func (c *Client) writePump() {
	const op = "delivery.http.ws_client.writePump"

	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				c.logger.Errorf("[%s]: failed to set write deadline: %v", op, err)
				return
			}
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					c.logger.Warnf("[%s]: failed to write close message: %v", op, err)
				}
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				c.logger.Errorf("[%s]: NextWriter error: %v", op, err)
				return
			}
			if _, err := w.Write(message); err != nil {
				c.logger.Errorf("[%s]: failed to write message: %v", op, err)
				return
			}

			n := len(c.send)
			for i := 0; i < n; i++ {
				if _, err := w.Write([]byte{'\n'}); err != nil {
					c.logger.Errorf("[%s]: failed to write delimiter: %v", op, err)
					return
				}
				if _, err := w.Write(<-c.send); err != nil {
					c.logger.Errorf("[%s]: failed to write buffered message: %v", op, err)
					return
				}
			}

			if err := w.Close(); err != nil {
				c.logger.Errorf("[%s]: writer close error: %v", op, err)
				return
			}

		case <-ticker.C:
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				c.logger.Errorf("[%s]: failed to set write deadline (ping): %v", op, err)
				return
			}
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.logger.Warnf("[%s]: ping failed: %v", op, err)
				return
			}
		}
	}
}
