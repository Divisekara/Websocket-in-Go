package main

import "github.com/gorilla/websocket"

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
}

func NewClient(connection *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: connection,
		manager:    manager,
	}
}

func (c *Client) readMessages() (int, []byte) {
	messageType, message, err := c.connection.ReadMessage()
	if err != nil {
		panic(err)
	}
	return messageType, message
}
