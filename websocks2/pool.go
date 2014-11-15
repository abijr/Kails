package websocks2

import (
	"log"
	"sync"
)

type pool struct {
	loggedIn map[string]*connection
	sync.RWMutex
}

var friends = &pool{
	loggedIn: make(map[string]*connection),
}

func (c *connection) stablishRTC(pc *connection) {
}

// m.Data["peer"], is the peer to chat with
// m.Data["request"], is the request type (videochat, or chat)
func (c *connection) bootstrapRTC(m *Message) {
	friends.RLock()
	pc, ok := friends.loggedIn[m.Data["peer"]]
	if !ok {
		return
	}
	friends.RUnlock()

	data := map[string]string{
		"type":   m.Data["request"],
		"webrtc": c.user.Webrtc,
		"user":   c.user.Name,
	}

	pc.ws.WriteJSON(data)
}

func (c *connection) register() {
	if c.registered {
		log.Println("Not registering user")
		return
	}

	// Lock the map for writing
	friends.Lock()
	defer friends.Unlock()

	// Register user
	friends.loggedIn[c.user.Name] = c
	c.registered = true
}

func (c *connection) unRegister() {
	// Lock the map for writing
	friends.Lock()
	defer friends.Unlock()

	// Register user
	if _, ok := friends.loggedIn[c.user.Name]; ok {
		delete(friends.loggedIn, c.user.Name)
	}
}
