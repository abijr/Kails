package routes

import (
	"log"

	"bitbucket.com/abijr/kails/middleware"
	"bitbucket.com/abijr/kails/models"
)

func Home(ctx *middleware.Context) {
	if ctx.IsLogged {
		ctx.Data["Title"] = "Home"
		ctx.Data["Name"] = ctx.User.Username
		ctx.Data["Email"] = ctx.User.Email
		ctx.Data["Since"] = ctx.User.Created.Format("January 21 of 2009")
		ctx.HTML(200, "user/home")
	} else {
		ctx.Data["Title"] = "Welcome"
		ctx.HTML(200, "main/main")
	}
}

func SignUp(ctx *middleware.Context) {
	ctx.Data["Title"] = "Sign Up"
	ctx.HTML(200, "user/signup")
}

func SignUpPost(ctx *middleware.Context, form models.UserSignupForm) {
	ctx.Data["Title"] = "Signed Up!"
	ctx.Data["Name"] = form.Username
	err := models.NewUser(form)
	if err != nil {
		ctx.HTML(501, "")
	}
	ctx.HTML(200, "user/signup")
}

func Login(ctx *middleware.Context) {
	if ctx.IsLogged {
		ctx.Redirect("/")
	} else {
		ctx.Data["Title"] = "Login"
		ctx.HTML(200, "user/login")
	}
}

func LoginPost(ctx *middleware.Context, form models.UserLoginForm) {
	ctx.Data["Title"] = "Home"
	user, err := models.UserByEmail(form.Email)
	if err != nil {
		// TODO: fill this up.
		log.Println(err)
		log.Println(user)
	}

	ctx.User = *user
	ctx.Session.Set("name", user.Username)
	ctx.IsLogged = true

	ctx.Redirect("/")
}
