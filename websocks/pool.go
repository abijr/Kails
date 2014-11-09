package websocks

import "log"

type pool struct {
	name    string
	spanish map[*connection]bool
	english map[*connection]bool

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

// Chat pool
var p = pool{
	name:       "chat",
	spanish:    make(map[*connection]bool),
	english:    make(map[*connection]bool),
	register:   make(chan *connection),
	unregister: make(chan *connection),
}

// Videochat pool
var v = pool{
	name:       "videochat",
	spanish:    make(map[*connection]bool),
	english:    make(map[*connection]bool),
	register:   make(chan *connection),
	unregister: make(chan *connection),
}

func (c *connection) stablishRTC(pc *connection) {
	c.ws.WriteJSON(pc.user)
	pc.ws.WriteJSON(c.user)
}

func (c *connection) register(p *pool) {
	log.Printf("Registering user `%v` in pool `%v`.\n", c.user.Name, p.name)
	var partnerPool, userPool map[*connection]bool
	if c.language == "english" {
		partnerPool = p.spanish
		userPool = p.english
	} else {
		partnerPool = p.english
		userPool = p.spanish
	}
	for partner := range partnerPool {
		c.stablishRTC(partner)
		return
	}

	userPool[c] = true
}

func (p *pool) run() {
	for {
		select {
		case c := <-p.register:
			c.register(p)
		case c := <-p.unregister:
			if c.language == "english" {
				delete(p.english, c)
			} else {
				delete(p.spanish, c)
			}
			close(c.send)
		}
	}
}
