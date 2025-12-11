package http

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestHub_Run(t *testing.T) {
	hub := NewHub()
	go hub.Run()

	trackID := uuid.New()

	client1 := &Client{
		hub:     hub,
		trackID: trackID,
		send:    make(chan []byte, 10),
		conn:    &websocket.Conn{},
	}

	client2 := &Client{
		hub:     hub,
		trackID: trackID,
		send:    make(chan []byte, 10),
		conn:    &websocket.Conn{},
	}

	otherTrackID := uuid.New()
	client3 := &Client{
		hub:     hub,
		trackID: otherTrackID,
		send:    make(chan []byte, 10),
		conn:    &websocket.Conn{},
	}

	hub.register <- client1
	hub.register <- client2
	hub.register <- client3

	time.Sleep(50 * time.Millisecond)

	assert.Len(t, hub.rooms[trackID], 2)
	assert.Len(t, hub.rooms[otherTrackID], 1)

	msgData := []byte("Hello World")
	hub.broadcast <- Message{
		TrackID: trackID,
		Data:    msgData,
	}

	time.Sleep(50 * time.Millisecond)

	select {
	case msg := <-client1.send:
		assert.Equal(t, msgData, msg)
	default:
		t.Error("Client 1 did not receive message")
	}

	select {
	case msg := <-client2.send:
		assert.Equal(t, msgData, msg)
	default:
		t.Error("Client 2 did not receive message")
	}

	select {
	case <-client3.send:
		t.Error("Client 3 received message meant for another room")
	default:
	}

	hub.unregister <- client1
	time.Sleep(50 * time.Millisecond)

	assert.Len(t, hub.rooms[trackID], 1)

	_, ok := <-client1.send
	assert.False(t, ok)

	hub.unregister <- client2
	time.Sleep(50 * time.Millisecond)

	_, exists := hub.rooms[trackID]
	assert.False(t, exists)
}
