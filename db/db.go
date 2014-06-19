package db

import "labix.org/v2/mgo"

var (
	db *mgo.Session
)

func init() {
	Init()
}

func Init() {
	var err error
	db, err = mgo.Dial("localhost:" + "27017")
	if err != nil {
		panic(err)
	}
}

func Close() {
	db.Close()
}

func Collection(collection string) *mgo.Collection {
	c := db.DB("kails").C(collection)
	return c
}
