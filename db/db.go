package db

import "github.com/diegogub/aranGO"

var (
	db *aranGO.Session
	DB *aranGO.Database
)

// TODO: there should be a better way to manage the database connection

// Init starts connection to mongodb server
func init() {
	var err error
	// TODO: Need to make this configurable somehow...
	db, err = aranGO.Connect("http://localhost:8529", "", "", false)
	if err != nil {
		panic(err)
	}

	DB = db.DB("kails")
}

// aranGO doesn't support the close operation
// // Close closes database connection
// func Close() {
// 	db.Close()
// }

// Collection returns a pointer to the named collection
func Collection(collection string) *aranGO.Collection {
	c := DB.Col(collection)
	return c
}
