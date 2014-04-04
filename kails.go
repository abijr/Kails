//wsrs is a simple spaced repetion web application.
package main

import (
	"bitbucket.com/abijr/kails/pool" // For handling auto-updatable templates
	//"labix.org/v2/mgo" // MongoDB handle
	//"labix.orr/v2/mgo/bson"
	"fmt"
	"github.com/go-martini/martini"
	// 	"log"
	"net/http"
	//"time"
)

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

	err := pool.Render("main", "main/index", values, rw)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
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
