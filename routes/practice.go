package routes

import (
	"log"
	"strconv"

	"github.com/go-martini/martini"

	"bitbucket.com/abijr/kails/middleware"
	"bitbucket.com/abijr/kails/models"
)

func Practice(ctx *middleware.Context) {
	ctx.Data["Title"] = "Practice"
	ctx.HTML(200, "practice/practice")
}

func Chat(ctx *middleware.Context) {
	ctx.Data["Title"] = "Chat"
	ctx.HTML(200, "practice/chat")
}

func Videochat(ctx *middleware.Context) {
	ctx.Data["Title"] = "Videochat"
	ctx.HTML(200, "practice/videochat")
}

func GetUser(ctx *middleware.Context, params martini.Params) {
	log.Println("User:", params["name"])
	user, err := models.UserByName(params["name"])

	if err == nil {
		log.Println(user.Lessons)
		ctx.JSON(200, user)
	}
}

func GetUserPrivileges(ctx *middleware.Context, params martini.Params) {
	level, err := strconv.Atoi(params["level"])

	if err != nil {
		log.Println("Error getting level", err)
	}

	privilege, err := models.GetUserPrivileges(level)

	if err == nil {
		log.Println("Friends: ", privilege.Friends)
		ctx.JSON(200, privilege)
	}
}
