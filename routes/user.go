package routes

import (
	"bitbucket.com/abijr/kails/middleware"
	"bitbucket.com/abijr/kails/models"
)

func Home(ctx *middleware.Context) {
	user, _ := models.UserByName("user1")
	ctx.Data["Name"] = user.Name
	ctx.Data["Title"] = "Welcome"
	ctx.HTML(200, "main/main")
}

func SignUp(ctx *middleware.Context) {
	ctx.Data["Title"] = "Sign Up"
	ctx.HTML(200, "user/signup")
}
