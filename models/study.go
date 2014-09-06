package models

import (
	"errors"

	"bitbucket.com/abijr/kails/db"

	"log"

	"labix.org/v2/mgo/bson"
)

// Level structure:
// Level ->
// 		Words ->
// 			Sentences

// Level contains level data
// it's structure allows data binding
// with database.
type Level struct {
	Id          int
	Description string
	Version     int
	Words       []Word
}

type Word struct {
	Word        string
	Level       int
	Class       string
	Translation string
	Sentences   []Sentence
}

type Sentence struct {
	Native      string
	Translation string
}

const (
	firstLevel = 1
)

// Utility variables
var (
	levels           = db.Collection("languages")
	errLevelNotExist = errors.New("level does not exist")
)

// LevelById finds the level by id and returns a pointer
// to the structure containing the data.
func LevelById(id int, lang string) (*Level, error) {
	var level *Level
	level = new(Level)

	// if smaller than first level, invalid level
	if id < firstLevel {
		return nil, errLevelNotExist
	}

	if lang == "" {
		return nil, errLevelNotExist
	}

	// Levels and words both exist in same collection
	// conflict prevented because only levels have `id` field
	err := levels.Find(bson.M{"id": id, "lang": lang}).One(level)
	log.Println("id", id, "lang", lang)
	if err != nil {
		return nil, err
	}
	return level, nil
}
