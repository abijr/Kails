package db

import "github.com/diegogub/aranGO"

var (
	db *aranGO.Session
)

func init() {
	Init()
}

// Init starts connection to mongodb server
func Init() {
	var err error
	// TODO: Need to make this configurable somehow...
	db, err = aranGO.Connect("http://localhost:8529", "", "", true)
	if err != nil {
		panic(err)
	}
}

// aranGO doesn't support the close operation
// // Close closes database connection
// func Close() {
// 	db.Close()
// }

// Collection returns a pointer to the named collection
func Collection(collection string) *aranGO.Collection {
	c := db.DB("kails").Col(collection)
	return c
}
