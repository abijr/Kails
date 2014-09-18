//Kails is a simple spaced repetion web application.
package main

import (
	"bitbucket.com/abijr/kails/db"
	"bitbucket.com/abijr/kails/middleware"
	"bitbucket.com/abijr/kails/models"
	"bitbucket.com/abijr/kails/routes"

	"github.com/abijr/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessions"
)

func main() {
	defer db.Close()
	// Set cookie store
	cookieStore := sessions.NewCookieStore([]byte("randomStuff"))
	m := martini.Classic()

	// Start the cookie handler
	m.Use(sessions.Sessions("kails_session", cookieStore))

	// Setup templates
	m.Use(render.Renderer(render.Options{
		Directory:            "templates",
		Languages:            []string{"en-US", "es-MX"},
		TranslationDirectory: "translations/all",
		Extensions:           []string{".tmpl.html"},
	}))

	m.Use(middleware.InitContext())
	m.Use(martini.Static("webapp/dist"))

	// Start the language handler
	// Serve the application on '/'
	m.Group("/webapp", func(r martini.Router) {

		m.Get("/study", routes.StudyPage)
		r.Get("/study/:id", routes.Study)
		r.Post("/study/:id", routes.StudyPost)

		r.Get("/settings", routes.Settings)
		r.Post("/settings", binding.Bind(routes.SettingsForm{}), routes.SettingsPost)
		m.Get("/program", routes.Program)
	})

	m.Get("/signup", middleware.Localizer, routes.SignUp)
	m.Post("/signup", middleware.Localizer, binding.Bind(models.UserSignupForm{}), routes.SignUpPost)

	m.Get("/login", middleware.Localizer, routes.Login)
	m.Post("/login", middleware.Localizer, binding.Bind(models.UserLoginForm{}), routes.LoginPost)

	m.Get("/logout", middleware.Localizer, routes.Logout)

	m.Get("/**", middleware.Localizer, routes.Home)

	// Launch server
	// It will automatically serve files under the "public" folder
	// public/css/file = localhost/css/file
	m.Run()
}
