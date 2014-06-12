//Kails is a simple spaced repetion web application.
package main

import (
	"log"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"bitbucket.com/abijr/kails/localization"

	"github.com/abijr/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

type Data struct {
	UserName string `user`
	Email    string `email`
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

	// Start the language handler
	// Serve the application on '/'
	m.Get("/", localizer, func(r render.Render, lang localization.Localizer) {
		result := Data{}
		err := c.Find(bson.M{"name": "user1"}).One(&result)
		if err != nil {
			panic(err)
		}
		result.UserName = "user1"
		log.Println(result)
		r.HTML(200, "main/main", result, lang.Get())
	})

	// Launch server
	// It will automatically serve files under the "public" folder
	// public/css/file = localhost/css/file
	m.Run()
}
