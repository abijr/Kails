//Kails is a simple spaced repetion web application.
package main

import (

	//"time"

	"math/rand"

	"bitbucket.com/abijr/kails/localization"

	// For handling auto-updatable templates

	//"labix.org/v2/mgo" // MongoDB handle
	//"labix.orr/v2/mgo/bson"

	"github.com/abijr/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
)

// func index(rw http.ResponseWriter, req *http.Request) {
// 	values := struct {
// 		Usuario    string
// 		Repasa     int
// 		Aprende    int
// 		PerRoots   int
// 		PerVocab   int
// 		PerGrammar int
// 	}{
// 		"potemkin",
// 		10,
// 		21,
// 		70,
// 		30,
// 		42,
// 	}
//
// 	if i := rand.Intn(2); i != 0 {
// 		log.Println(i)
//
// 		err := p.Render("main", "main", "es-MX", values, rw)
// 		if err != nil {
// 			log.Println(err)
// 		}
//
// 	} else {
// 		log.Println(i)
//
// 		err := p.Render("main", "main", "en-US", values, rw)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}
// }

func main() {
	rand.Seed(1024)

	// Set cookie store
	cookieStore := sessions.NewCookieStore([]byte("randomStuff"))
	localizer := localization.NewLocalizer(localization.Options{
		DefaultLanguage: "en-US",
	})
	/*
		session, err := mgo.Dial("localhost:"+dbPort)
		if err != nil {
			panic(err)
		}

		defer session.Close()
		c := session.DB("wsrs").C("users")

		if err != nil {
			panic(err)
		}
	*/
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
		r.HTML(200, "main/main", nil, lang.Get())
	})

	// Launch server
	// It will automatically serve files under the "public" folder
	// public/css/file = localhost/css/file
	m.Run()
}
