//Kails is a simple spaced repetion web application.
package main

import (
	"net/http"
	//"time"
	"log"
	"math/rand"

	"bitbucket.com/abijr/kails/pool" // For handling auto-updatable templates

	//"labix.org/v2/mgo" // MongoDB handle
	//"labix.orr/v2/mgo/bson"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessions"
	"github.com/nicksnyder/go-i18n/i18n/locale"
)

var p *pool.Pool

func index(rw http.ResponseWriter, req *http.Request) {
	values := struct {
		Usuario    string
		Repasa     int
		Aprende    int
		PerRoots   int
		PerVocab   int
		PerGrammar int
	}{
		"potemkin",
		10,
		21,
		70,
		30,
		42,
	}

	if i := rand.Intn(2); i != 0 {
		log.Println(i)

		err := p.Render("main", "main", "es-MX", values, rw)
		if err != nil {
			log.Println(err)
		}

	} else {
		log.Println(i)

		err := p.Render("main", "main", "en-US", values, rw)
		if err != nil {
			log.Println(err)
		}

	}

}

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
		if reqLang != "" {
			language, err = locale.New(reqLang)
			if err != nil {
				log.Println("Invalid locale, falling back to default: ", DEFAULT_LANGUAGE)
				language, _ = locale.New(DEFAULT_LANGUAGE)
			}
		}
		session.Set("Language", language.ID)
		log.Println("Accepted language recieved is: ", language.ID)
	} else {
		log.Println("Language loaded from session")
	}
}

func main() {
	rand.Seed(1024)

	// Setup templates
	p, _ = pool.NewPool("templates", "translations", []string{"en-US", "es-MX"})

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

	// Start the language handler
	// Serve the application on '/'
	m.Get("/", setLanguage, index)

	// Launch server
	// It will automatically serve files under the "public" folder
	// public/css/file = localhost/css/file
	m.Run()
}
