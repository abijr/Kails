package db

import "labix.org/v2/mgo"

var (
	db *mgo.Session
)

func init() {
	Init()
}

// Init starts connection to mongodb server
func Init() {
	var err error
	// Need to make this configurable somehow...
	db, err = mgo.Dial("localhost:" + "27017")
	if err != nil {
		panic(err)
	}
}

// Close closes database connection
func Close() {
	db.Close()
}

// Collection returns a pointer to the named collection
func Collection(collection string) *mgo.Collection {
	c := db.DB("kails").C(collection)
	return c
}
