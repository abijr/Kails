package routes

import (
	//"log"

	//"github.com/go-martini/martini"

	"bitbucket.com/abijr/kails/middleware"
	//"bitbucket.com/abijr/kails/models"
)

func Practice(ctx *middleware.Context) {
	ctx.Data["Title"] = "Practice"
	ctx.HTML(200, "practice/main")
}

func Chat(ctx *middleware.Context) {
	ctx.Data["Title"] = "Chat"
	ctx.HTML(200, "practice/chat")
}

func Videochat(ctx *middleware.Context) {
	ctx.Data["Title"] = "Videochat"
	ctx.HTML(200, "practice/videochat")
}