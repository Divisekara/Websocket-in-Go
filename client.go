package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
	}
}

func (c *Client) readMessages() {
	defer func() {
		// cleanup connection
		c.manager.removeClient(c)
	}()

	for {
		// messageTypes can be ping,pong, data, binary messages. We only use either binary or text data
		messageType, payload, err := c.connection.ReadMessage()
		if err != nil {
			// connection can be unexpectedly close or slow
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// we dont need to log when connection is closed expectedly by user
				log.Printf("error reading message: %v", err)
			}
			break
		}
		log.Println(messageType)
		log.Println(string(payload))
		_ = c.connection.WriteMessage(1, []byte("asitha"))
	}
}
