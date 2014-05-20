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
	// 	"fmt"
	"github.com/go-martini/martini"
	// 	"log"
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

func main() {
	rand.Seed(1024)
	p, _ = pool.NewPool("templates", "translations", []string{"en-US", "es-MX"})
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

	// Serve the application on '/'
	m.Get("/", index)

	// Launch server
	// It will automatically serve files under the "public" folder
	// public/css/file = localhost/css/file
	m.Run()
}
