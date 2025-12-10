package http

import (
	"github.com/google/uuid"
)

type Message struct {
	TrackID uuid.UUID
	Data    []byte
}

type Hub struct {
	rooms      map[uuid.UUID]map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		rooms:      make(map[uuid.UUID]map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			if h.rooms[client.trackID] == nil {
				h.rooms[client.trackID] = make(map[*Client]bool)
			}
			h.rooms[client.trackID][client] = true

		case client := <-h.unregister:
			if clients, ok := h.rooms[client.trackID]; ok {
				if _, ok := clients[client]; ok {
					delete(clients, client)
					close(client.send)

					if len(clients) == 0 {
						delete(h.rooms, client.trackID)
					}
				}
			}

		case message := <-h.broadcast:
			if clients, ok := h.rooms[message.TrackID]; ok {
				for client := range clients {
					select {
					case client.send <- message.Data:
					default:
						close(client.send)
						delete(clients, client)
					}
				}
			}
		}
	}
}
