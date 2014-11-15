package websocks2

import (
	"log"

	"bitbucket.com/abijr/kails/middleware"
)

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

func ServeConnectedFriends(ctx *middleware.Context) {
	list, err := ctx.User.ListFriends()
	if err != nil {
		ctx.HTML(500, "error/error")
	}

	friends.RLock()
	defer friends.RUnlock()

	connected := make(map[string]bool, len(list))

	for _, f := range list {
		if _, ok := friends.loggedIn[f.Username]; ok {
			connected[f.Username] = true
		}
	}

	ctx.JSON(200, connected)
}
