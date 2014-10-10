package routes

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"bitbucket.com/abijr/kails/middleware"
	"bitbucket.com/abijr/kails/models"
	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2/bson"
)

type Card struct {
	Sentence models.Sentence
	Word     string
}

// Study returns the given lesson json definition
func Study(ctx *middleware.Context, params martini.Params) {
	// TODO: Add a JSON return error

	// Commented for testing....
	// if !ctx.IsLogged {
	// 	ctx.Redirect("/login")
	//	return
	// }

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//TODO: fill this up
		log.Println("couldn't parse int: ", params["id"])
		return
	}

	//----------------------------------
	// Just here for testing...			   |
	if ctx.User.StudyLanguage == "" { // 	   |
		ctx.User.StudyLanguage = "english" //  |
	} // 								   |
	// ---------------------------------

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
			card := Card{sentence, word.Word}
			session = append(session, card)
		}
	}

	ctx.JSON(200, session)
}

func StudyPost(ctx *middleware.Context, params martini.Params) {

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//TODO: fill this up
		log.Println("couldn't parse int: ", params["id"])
		return
	}

	// fmt.Println("hi")
	// io.Copy(os.Stdout, ctx.Req.Body)
	// fmt.Println("hi")

	decoder := json.NewDecoder(ctx.Req.Body)
	test := struct {
		Pass      bool `json:"pass"`
		Sentences []struct {
			Number int    `json:"number"`
			Score  int    `json:"score"`
			Word   string `json:"word"`
		} `json:",omitempty"`
	}{
		Pass: false,
	}
	err = decoder.Decode(&test)
	log.Println(test)
	if err != nil {
		// TODO: Need to remove this and properly handle the error
		log.Println("couldn't decode request body:", err)
		return
	}

	if test.Pass {
		experienceGained := 15
		ctx.User.AddExperience(experienceGained)
		newUserLesson := models.UserLesson{id, time.Now()}

		err := ctx.User.UpdateLesson(newUserLesson)
		if err != nil {
			log.Println("couldn't update user info:", err)
		}

		ctx.JSON(200, bson.M{
			"ExperienceGained": experienceGained,
		})
	}

}

func StudyPage(ctx *middleware.Context) {
	ctx.HTML(200, "study/study")
}

func Program(ctx *middleware.Context) {
	p, err := models.ProgramByLanguage(ctx.User.StudyLanguage)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.Data["Lessons"] = p.Lessons
	ctx.Data["Experience"] = ctx.User.Experience
	ctx.Data["Level"] = ctx.User.Level
	ctx.Data["NextLevel"] = ctx.User.Level + 1
	ctx.Data["PercentDone"] = ctx.User.PercentToNextLevel()
	ctx.HTML(200, "user/program")
}
