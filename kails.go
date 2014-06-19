//Kails is a simple spaced repetion web application.
package main

import (
	"log"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	_ "bitbucket.com/abijr/kails/db"
	"bitbucket.com/abijr/kails/middleware"
	"bitbucket.com/abijr/kails/middleware/localization"

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

	db, err := mgo.Dial("localhost:" + "27017")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	c := db.DB("kails").C("users")

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
		result := Data{}
		err := c.Find(bson.M{"name": "user1"}).One(&result)
		if err != nil {
			panic(err)
		}
		result.Name = "user1"
		log.Println(result)
		ctx.HTML(200, "main/main")
	})

	// Launch server
	// It will automatically serve files under the "public" folder
	// public/css/file = localhost/css/file
	m.Run()
}
