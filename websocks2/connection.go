// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocks2

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type info struct {
	Name   string `json:"name"`
	Webrtc string `json:"webrtc"`
}

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	user info

	language string

	connectionType string

	registered bool
}

type Message struct {
	Type string
	Data map[string]string
}

// readPump pumps messages from the websocket connection to the hub.
func (c *connection) readPump() {
	message := new(Message)

	defer func() {
		log.Printf("Unregistering user `%v` with webrtc key `%v`", c.user.Name, c.user.Webrtc)
		c.unRegister()
	}()

	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	c.register()

	for {
		err := c.ws.ReadJSON(message)
		if err != nil {
			log.Println("##### error reading message")
			break
		}
		log.Printf("Message: %v", message)

		c.user.Webrtc = message.Data["id"]
		c.connectionType = message.Type

		switch message.Type {
		case "request":
			c.bootstrapRTC(message)
			continue
		default:
			continue
		}
	}
}

// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (c *connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
