//Kails is a simple spaced repetion web application.
package main

import (
	"bitbucket.com/abijr/kails/db"
	"bitbucket.com/abijr/kails/middleware"
	"bitbucket.com/abijr/kails/routes"

	"github.com/abijr/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

func main() {
	db.Init()
	defer db.Close()
	// Set cookie store
	cookieStore := sessions.NewCookieStore([]byte("randomStuff"))
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
	m.Get("/", middleware.Localizer, routes.Home)

	// Launch server
	// It will automatically serve files under the "public" folder
	// public/css/file = localhost/css/file
	m.Run()
}
