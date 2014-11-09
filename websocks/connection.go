// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocks

import (
	"fmt"
	"log"
	"time"

	"bitbucket.com/abijr/kails/middleware"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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
}

type Message struct {
	Type string
	Data string
}

// readPump pumps messages from the websocket connection to the hub.
func (c *connection) readPump() {
	defer func() {
		log.Println("Unregistering.................")
		p.unregister <- c
	}()
	var registered bool
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var message Message
		err := c.ws.ReadJSON(&message)
		if err != nil {
			break
		}
		c.user.Webrtc = message.Data
		fmt.Printf("Message: %v\n", c.user.Webrtc)
		if !registered {
			if message.Type == "chat" {
				p.register <- c
				registered = true
			} else if message.Type == "videochat" {
				v.register <- c
				registered = true
			}
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

// ServeWs handles websocket requests from the peer.
func ServeWs(ctx *middleware.Context) {
	log.Println("Serving websockets upgrader to user: ", ctx.User.Username)
	ws, err := upgrader.Upgrade(ctx.Res, ctx.Req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	c := &connection{
		send:     make(chan []byte, 256),
		ws:       ws,
		language: ctx.User.StudyLanguage,
		user:     info{Name: ctx.User.Username},
	}
	go c.writePump()
	c.readPump()
}
