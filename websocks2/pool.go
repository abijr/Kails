package websocks2

import (
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

func (c *connection) register() {
	// Lock the map for writing
	friends.Lock()
	defer friends.Unlock()

	// Register user
	friends.loggedIn[c.user.Name] = c
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
