package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

/**
 * File: hub.go
 * Date: 2021-08-13 17:25:44
 * Creator: Sean Patrick Hagen <sean.hagen@gmail.com>
 */

// shamelessly copied from https://github.com/gorilla/websocket/blob/master/examples/chat/hub.go

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	ticker := time.NewTicker(time.Second * 10)

	go func() {
		for {
			select {
			case tick := <-ticker.C:
				log.Printf("tick: %v", tick.String())
				msg := fmt.Sprintf("ping %v", tick.String())
				h.broadcast <- []byte(msg)
				log.Printf("broadcast tick message done")
			case client := <-h.register:
				log.Printf("Registering client %v", client.id)
				h.clients[client] = true
				hiMsg := fmt.Sprintf("hello client %v", client.id)
				client.sendMsg(string(h.msgToJson([]byte(hiMsg))))
			case client := <-h.unregister:
				log.Printf("Client %v unregistering", client.id)
				if _, ok := h.clients[client]; ok {
					delete(h.clients, client)
					close(client.send)
				}
			default:
			}
		}
	}()

	for {
		select {
		case message := <-h.broadcast:
			m := h.msgToJson(message)
			fmt.Printf("sending message: \n\t'%v'\n", string(message))
			for client := range h.clients {
				log.Printf("Sending to client %v", client.id)
				select {
				case client.send <- m:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// msgToJson ...
func (h Hub) msgToJson(in []byte) []byte {
	m := string(in)
	msg := struct {
		M string `json:"m"`
	}{M: m}

	b, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Unable to convert string to json: %v", err)
		o := "{}"
		return []byte(o)
	}
	return b
}
