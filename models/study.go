package models

import (
	"errors"

	"bitbucket.com/abijr/kails/db"

	"labix.org/v2/mgo/bson"
)

type Level struct {
	Id          int
	Description string
	Version     int
	Words       []Term
}

type Term struct {
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

func LevelById(id int, lang string) (*Level, error) {
	var level Level
	level = new(Level)

	// if smaller than first level, invalid level
	if id < firstLevel {
		return nil, errLevelNotExist
	}

	// Levels and words both exist in same collection
	// conflict prevented because only levels have `id` field
	err := levels.Find(bson.M{"id": id, "lang": lang}).One(level)
	if err != nil {
		return nil, err
	}
	return user, nil
}
