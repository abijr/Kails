//wsrs is a simple spaced repetion web application.
package main

import (
	"bitbucket.com/abijr/wsrs/pool"
	//"labix.org/v2/mgo"
	//"labix.orr/v2/mgo/bson"
	"fmt"
	"log"
	"net/http"
	//"time"
	"github.com/gorilla/mux"
)

func index(rw http.ResponseWriter, req *http.Request) {
	values := struct {
		Repasa     int
		Aprende    int
		PerRoots   int
		PerVocab   int
		PerGrammar int
	}{
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
		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}

		defer session.Close()
		c := session.DB("wsrs").C("users")

		if err != nil {
			panic(err)
		}
	*/

	r := mux.NewRouter()

	// Serve the /assets directory and its contents
	r.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	// Serve the application on '/'
	r.HandleFunc("/", index)

	// Handle the routed stuff
	http.Handle("/", r)

	port := "3000"

	// Launch server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
