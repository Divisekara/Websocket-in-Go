package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var (
	webSocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	client ClientList
	sync.RWMutex
}

// NewManager factory method for Manger struct
func NewManager(client ClientList) *Manager {
	return &Manager{client: make(ClientList)}
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection")

	// upgrade regular http connection into websocket
	conn, err := webSocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, m)
	m.addClient(client)

}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	m.client[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.client[client]; ok {
		client.connection.Close()
		delete(m.client, client)
	}
}
