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