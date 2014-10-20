package routes

import (
	"encoding/json"
	"log"
	"math"
	"strconv"
	"time"

	"bitbucket.com/abijr/kails/middleware"
	"bitbucket.com/abijr/kails/models"
	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2/bson"
)

// Card is the structure of the objects in the array
// array  sent to the webapp when it requests a lesson.
type Card struct {
	Sentence   models.Sentence
	Word       string
	Definition string
}

// Study returns the given lesson json definition
func Study(ctx *middleware.Context, params martini.Params) {
	// TODO: Add a JSON return error

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//TODO: fill this up
		log.Println("couldn't parse int: ", params["id"])
		return
	}

	lesson, err := models.LessonById(id, ctx.User.StudyLanguage)
	if err != nil {
		//TODO: fill this up
		log.Println("Err getting lesson: ", err)
		return
	}

	// lesson is made up of cards
	session := make([]Card, 0, len(lesson.Words)*3)

	for _, word := range lesson.Words {
		for _, sentence := range word.Sentences {
			card := Card{sentence, word.Word, word.Definition}
			session = append(session, card)
		}
	}

	ctx.JSON(200, session)
}

// StudyPost parses results of the study session.
func StudyPost(ctx *middleware.Context, params martini.Params) {

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//TODO: fill this up
		log.Println("couldn't parse int: ", params["id"])
		return
	}

	// Parses JSON into `test` struct
	decoder := json.NewDecoder(ctx.Req.Body)
	test := struct {
		Pass       bool     `json:"Pass"`
		WrongWords []string `json:"WrongWords,omitempty"`
	}{
		Pass: false,
	}
	err = decoder.Decode(&test)
	if err != nil {
		// TODO: Need to remove this and properly handle the error
		log.Println("couldn't decode request body:", err)
		return
	}

	// If the lesson was completed succesfully
	// TODO: Logic can be simplified here.
	if test.Pass {
		experienceGained := 15
		ctx.User.AddExperience(experienceGained)
		newUserLesson := models.UserLesson{id, time.Now(), 0}

		err := ctx.User.UpdateLesson(newUserLesson)
		if err != nil {
			log.Println("couldn't update user info:", err)
		}

		ctx.JSON(200, bson.M{
			"ExperienceGained": experienceGained,
		})
	} else {
		ctx.JSON(200, bson.M{"ExperienceGained": 0})
	}

}

// StudyPage returns empty html for the webapp to fill
func StudyPage(ctx *middleware.Context) {
	ctx.HTML(200, "study/study")
}

func Flashcard(ctx *middleware.Context) {
	ctx.HTML(200, "study/flashcard")
}

// TODO: Refactor db structure so that 'UserLessons' have the description
// embedded
type lesson struct {
	Id          int
	Description string
	Opacity     float64
}

// Program returns the main page with the users lessons,
// level, and experience points.
func Program(ctx *middleware.Context) {
	p, err := models.ProgramByLanguage(ctx.User.StudyLanguage)
	if err != nil {
		log.Println(err)
		return
	}

	studyLessons := make([]*lesson, len(ctx.User.Lessons))
	for i := range studyLessons {

		op := 1.0
		days := time.Since(ctx.User.Lessons[strconv.Itoa(p.Lessons[i].Id)].LastReview).Hours() / 24
		if days < 10 {
			op -= .5 / math.Ceil(days)
		}
		studyLessons[i] = &lesson{p.Lessons[i].Id, p.Lessons[i].Description, op}
	}
	ctx.Data["Lessons"] = studyLessons
	ctx.Data["Experience"] = ctx.User.Experience
	ctx.Data["Level"] = ctx.User.Level
	ctx.Data["NextLevel"] = ctx.User.Level + 1
	ctx.Data["PercentDone"] = ctx.User.PercentToNextLevel()
	ctx.HTML(200, "user/program")
}
