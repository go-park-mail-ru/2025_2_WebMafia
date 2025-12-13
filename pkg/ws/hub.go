package ws

import (
	"context"
)

type Message struct {
	Topic string
	Data  []byte
}

type Hub struct {
	rooms      map[string]map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		rooms:      make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Run(ctx context.Context) {
	defer func() {
		for _, clients := range h.rooms {
			for client := range clients {
				close(client.send)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return

		case client := <-h.register:
			if h.rooms[client.topic] == nil {
				h.rooms[client.topic] = make(map[*Client]bool)
			}
			h.rooms[client.topic][client] = true

		case client := <-h.unregister:
			if clients, ok := h.rooms[client.topic]; ok {
				if _, ok := clients[client]; ok {
					delete(clients, client)
					close(client.send)
					if len(clients) == 0 {
						delete(h.rooms, client.topic)
					}
				}
			}

		case message := <-h.broadcast:
			if clients, ok := h.rooms[message.Topic]; ok {
				for client := range clients {
					select {
					case client.send <- message.Data:
					default:
						delete(clients, client)
						close(client.send)
						if len(clients) == 0 {
							delete(h.rooms, message.Topic)
						}
					}
				}
			}
		}
	}
}

func (h *Hub) Register(client *Client) {
	h.register <- client
}

func (h *Hub) Unregister(client *Client) {
	h.unregister <- client
}

func (h *Hub) BroadcastTo(topic string, data []byte) {
	h.broadcast <- Message{
		Topic: topic,
		Data:  data,
	}
}
