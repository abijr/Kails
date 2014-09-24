package routes

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"bitbucket.com/abijr/kails/middleware"
	"bitbucket.com/abijr/kails/models"
	"github.com/go-martini/martini"
)

type Card struct {
	Sentence models.Sentence
	Word     string
}

// Study returns the given level json definition
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

	level, err := models.LevelById(id, ctx.User.StudyLanguage)
	if err != nil {
		//TODO: fill this up
		log.Println("Err getting level: ", err)
		return
	}

	// lesson is made up of cards
	lesson := make([]Card, 0, len(level.Words)*3)

	for _, word := range level.Words {
		for _, sentence := range word.Sentences {
			card := Card{sentence, word.Word}
			lesson = append(lesson, card)
		}
	}

	ctx.JSON(200, lesson)

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
		newUserLevel := models.UserLevel{id, time.Now()}

		err := ctx.User.UpdateLevel(newUserLevel)
		if err != nil {
			log.Println("couldn't update user info:", err)
		}
	}

}

func Program(ctx *middleware.Context) {
	p, err := models.ProgramByLanguage(ctx.User.StudyLanguage)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.Data["Levels"] = p.Levels
	ctx.HTML(200, "user/program")
}

func StudyPage(ctx *middleware.Context) {
	ctx.HTML(200, "study/study")
}
