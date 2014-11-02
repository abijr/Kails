package websocks

import "log"

type pool struct {
	spanish map[*connection]bool
	english map[*connection]bool

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

var p = pool{
	spanish:    make(map[*connection]bool),
	english:    make(map[*connection]bool),
	register:   make(chan *connection),
	unregister: make(chan *connection),
}

func (c *connection) stablishRTC(pc *connection) {
	c.ws.WriteJSON(pc.user)
	pc.ws.WriteJSON(c.user)
}

func (c *connection) register() {
	log.Printf("Registering user `%v` in pool.\n", c.user.Name)
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
			c.register()
		case c := <-p.unregister:
			log.Println("Unregistering with method1: ", c.user.Name)
			if c.language == "english" {
				delete(p.english, c)
			} else {
				delete(p.spanish, c)
			}
			close(c.send)
			log.Println(p.english)
			log.Println(p.spanish)
			log.Println(".............................................")
		}
	}
}
