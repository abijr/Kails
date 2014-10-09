package models

import (
	"errors"
	"log"

	"bitbucket.com/abijr/kails/db"

	"gopkg.in/mgo.v2/bson"
)

const (
	wordType    = "program"
	lessonType  = "lesson"
	programType = "program"
)

const (
	firstLesson = 1
)

// Utility variables
var (
	languages           = db.Collection("languages")
	errLessonNotExist   = errors.New("lesson does not exist")
	errLanguageNotExist = errors.New("language does not exist")
)

// Lesson structure:
// Lesson ->
// 		Words ->
// 			Sentences

// Lesson contains lesson data
// it's structure allows data binding
// with database.
type Lesson struct {
	Id          int
	Description string
	Version     int
	Words       []Word
}

type Word struct {
	Word        string
	Lesson      int
	Class       string
	Translation string
	Sentences   []Sentence
}

type Sentence struct {
	Native      string
	Translation string
}

// LessonById finds the lesson by id and returns a pointer
// to the structure containing the data.
func LessonById(id int, lang string) (*Lesson, error) {
	var lesson *Lesson
	lesson = new(Lesson)

	// if smaller than first lesson, invalid lesson
	if id < firstLesson {
		return nil, errLessonNotExist
	}

	if lang == "" {
		return nil, errLessonNotExist
	}

	err := languages.First(bson.M{"Id": id, "Language": lang, "Type": lessonType}, lesson)
	if err != nil {
		return nil, err
	}

	return lesson, nil
}

type Program struct {
	Language string       "Language"
	Lessons  []LessonInfo "Lessons"
}

type LessonInfo struct {
	Id          int    "Id"
	Description string "Description"
}

func ProgramByLanguage(lang string) (*Program, error) {
	var program *Program
	program = new(Program)

	if lang == "" {
		return nil, errLanguageNotExist
	}

	err := languages.First(bson.M{"Language": lang, "Type": programType}, program)
	if err != nil {
		return nil, err
	}

	log.Println(program)
	return program, nil
}
