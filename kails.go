//Kails is a simple spaced repetion web application.
package main

import (
	"net/http"
	//"time"
	"log"
	"math/rand"

	// For handling auto-updatable templates

	//"labix.org/v2/mgo" // MongoDB handle
	//"labix.orr/v2/mgo/bson"

	"github.com/abijr/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"github.com/nicksnyder/go-i18n/i18n/locale"
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

const (
	DEFAULT_LANGUAGE = "en-US"
)

func setLanguage(req *http.Request, session sessions.Session) {
	// Get language from session
	sessLang := session.Get("Language")

	// check if it isn't set
	if _, ok := sessLang.(string); !ok {
		var language *locale.Locale
		var err error

		// get language from http header
		reqLang := req.Header.Get("Accept-Language")
		language, err = locale.New(reqLang)
		if err != nil {
			log.Printf("Error: %v, falling back to default language:  %v", err, DEFAULT_LANGUAGE)
			language, _ = locale.New(DEFAULT_LANGUAGE)
		}
		session.Set("Language", language.ID)
		log.Println("Accepted language recieved is: ", language.ID)
	} else {
		log.Println("Language loaded from session")
	}
}

func main() {
	rand.Seed(1024)

	// Set cookie store
	cookieStore := sessions.NewCookieStore([]byte("randomStuff"))
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
	m.Get("/", setLanguage, func(r render.Render) {
		r.HTML(200, "main/main", nil, "en-US")
	})

	// Launch server
	// It will automatically serve files under the "public" folder
	// public/css/file = localhost/css/file
	m.Run()
}
