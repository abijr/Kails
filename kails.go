//Kails is a simple spaced repetion web application.
package main

import (
	_ "bitbucket.com/abijr/kails/db"
	"bitbucket.com/abijr/kails/middleware"
	"bitbucket.com/abijr/kails/middleware/localization"
	"bitbucket.com/abijr/kails/models"

	"github.com/abijr/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

// Data represents the object of a user
type Data struct {
	Name     string `name`
	Email    string `email`
	Password string `pass`
}

func main() {
	// Set cookie store
	cookieStore := sessions.NewCookieStore([]byte("randomStuff"))
	// Set up localizer middleware
	localizer := localization.NewLocalizer(localization.Options{
		DefaultLanguage: "en-US",
	})
	m := martini.Classic()

	// Start the cookie handler
	m.Use(sessions.Sessions("kails_session", cookieStore))

	// Setup templates
	m.Use(render.Renderer(render.Options{
		Directory:            "templates",
		Languages:            []string{"en-US"},
		TranslationDirectory: "translations/all",
		Extensions:           []string{".tmpl.html"},
	}))

	m.Use(middleware.InitContext())

	// Start the language handler
	// Serve the application on '/'
	m.Get("/", localizer, func(ctx *middleware.Context) {
		user, _ := models.UserByName("user1")
		ctx.Data["Name"] = user.Name
		ctx.Data["Title"] = "Welcome"
		ctx.HTML(200, "main/main")
	})

	// Launch server
	// It will automatically serve files under the "public" folder
	// public/css/file = localhost/css/file
	m.Run()
}
