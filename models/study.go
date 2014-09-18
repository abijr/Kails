package models

import (
	"errors"

	"bitbucket.com/abijr/kails/db"

	"log"

	"gopkg.in/mgo.v2/bson"
)

const (
	wordType    = "program"
	levelType   = "level"
	programType = "program"
)

const (
	firstLevel = 1
)

// Utility variables
var (
	languages           = db.Collection("languages")
	errLevelNotExist    = errors.New("level does not exist")
	errLanguageNotExist = errors.New("language does not exist")
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

	err := languages.Find(bson.M{"id": id, "lang": lang, "type": levelType}).One(level)
	log.Println("id", id, "lang", lang)
	if err != nil {
		return nil, err
	}
	return level, nil
}

type Program struct {
	Language string      "lang"
	Levels   []LevelInfo "levels"
}

type LevelInfo struct {
	Id          int    "id"
	Description string "desc"
}

func ProgramByLanguage(lang string) (*Program, error) {
	var program *Program
	program = new(Program)

	if lang == "" {
		return nil, errLanguageNotExist
	}

	err := languages.Find(bson.M{"lang": lang, "type": programType}).One(program)
	log.Println(program)
	if err != nil {
		return nil, err
	}
	return program, nil
}
