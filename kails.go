//Kails is a simple spaced repetition web application.
package main

import (
	_ "bitbucket.com/abijr/kails/db"
	"bitbucket.com/abijr/kails/middleware"
	"bitbucket.com/abijr/kails/models"
	"bitbucket.com/abijr/kails/routes"
	"bitbucket.com/abijr/kails/websocks"

	"github.com/abijr/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessions"
)

func main() {

	// Set cookie store
	cookieStore := sessions.NewCookieStore([]byte("randomStuff"))
	m := martini.Classic()

	// Start the cookie handler
	m.Use(sessions.Sessions("kails_session", cookieStore))

	// Setup templates
	m.Use(render.Renderer(render.Options{
		Directory:            "templates",
		Languages:            []string{"en-us", "es-mx"},
		TranslationDirectory: "translations/all",
		Extensions:           []string{".tmpl.html"},
	}))

	m.Use(middleware.InitContext())
	m.Use(middleware.Localizer)
	m.Use(martini.Static("webapp/dist"))

	// Start the language handler
	// Serve the application on '/'
	m.Group("/webapp", func(r martini.Router) {

		m.Get("/study", routes.StudyPage)
		r.Get("/study/:id", routes.Study)
		r.Post("/study/:id", routes.StudyPost)

		m.Get("/words", routes.WordsPage)
		m.Get("/words/all", routes.Words)
		// m.Post("/words", routes.WordsPost)

		// Muestra los ajustes actuales y permite editarlos
		r.Get("/settings", routes.Settings)
		r.Post("/settings", binding.Bind(routes.SettingsForm{}), routes.SettingsPost)

		m.Get("/program", routes.Program)

		// Busqueda de usuarios (con "prefix" search)
		m.Get("/search/:name", routes.UserSearch)

		// Obtiene la informacion del usuario
		m.Get("/user/:name", routes.UserPage)

		// Chat
		m.Get("/chat", routes.Chat)

		// Practica
		m.Get("/practice", routes.Practice)
		m.Get("/practice/:name", routes.GetUser)
		m.Post("/practice/:name", routes.AddTopic)

		m.Get("/videochat", routes.Videochat)
		m.Get("/friends", routes.Friends)
		m.Get("/friends/connected", routes.GetFriendsConnected)
		m.Get("/friends/:user", routes.GetFriends)
		m.Get("/friends/:user/:topic", routes.CheckFriendStatus)

		m.Get("/flashcard", routes.Flashcard)
	})

	m.Get("/signup", routes.SignUp)
	m.Post("/signup", binding.Bind(models.UserSignupForm{}), routes.SignUpPost)

	m.Get("/login", routes.Login)
	m.Post("/login", binding.Bind(models.UserLoginForm{}), routes.LoginPost)

	m.Get("/logout", routes.Logout)

	m.Get("/ws", websocks.ServeWs)

	// Default route, should be the last one loaded
	// Returns the angular main page so that if the
	// route's managed by angular it gets handled by it
	m.Get("/**", routes.Home)

	// Launch server
	// It will automatically serve files under the "public" folder
	// public/css/file = localhost/css/file
	m.Run()
}
