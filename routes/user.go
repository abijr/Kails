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

func Settings(ctx *middleware.Context) {
	ctx.Data["Title"] = "Settings"
	ctx.Data["currentLanguage"] = ctx.User.StudyLanguage
	var lang string
	if ctx.User.StudyLanguage == "english" {
		lang = "spanish"
	} else {
		lang = "english"
	}
	ctx.Data["otherLanguage"] = lang
	ctx.HTML(200, "user/settings")
}

type SettingsForm struct {
	Language string `form:"language"`
}

func SettingsPost(ctx *middleware.Context, form SettingsForm) {
	if form.Language == ctx.User.StudyLanguage {
		ctx.Redirect("/")
		return
	}

	err := ctx.User.UserUpdateStudyLanguage(form.Language)
	if err != nil {
		//TODO: put something here
	}

}

func SignUp(ctx *middleware.Context) {
	ctx.Data["Title"] = "Sign Up"
	ctx.HTML(200, "user/signup")
}

func SignUpPost(ctx *middleware.Context, form models.UserSignupForm) {
	err := models.NewUser(form)
	if err != nil {
		log.Println(err)
		ctx.HTML(501, "")
		return
	}
	ctx.Redirect("/")
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

func Logout(ctx *middleware.Context) {
	if ctx.IsLogged {
		// This is necessary
		ctx.Session.Clear()

		// This is for making sure
		ctx.User = models.User{} // blank user
		ctx.IsLogged = false

		log.Println("user logged out")
	}

	ctx.Redirect("/")
}
