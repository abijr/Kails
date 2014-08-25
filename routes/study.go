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

// TODO: Add a JSON return error

func Study(ctx *middleware.Context, params martini.Params) {
	if !ctx.IsLogged {
		ctx.Redirect("/login")
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//TODO: fill this up
		log.Println("couldn't parse int: ", params["id"])
		return
	}

	level, err := models.LevelById(id, ctx.User.StudyLang)
	if err != nil {
		//TODO: fill this up
		log.Println("Err getting level: ", err)
		return
	}

	ctx.JSON(200, level)

}

func StudyResponse(ctx *middleware.Context) {

	// fmt.Println("hi")
	// io.Copy(os.Stdout, ctx.Req.Body)
	// fmt.Println("hi")

	decoder := json.NewDecoder(ctx.Req.Body)
	test := struct {
		Pass  bool `json:"pass"`
		Level int  `json:"level"`
		// Sentences []struct {
		// 	Number int
		// 	Score  int
		// 	Word   string
		// } `json:",omitempty"`
	}{
		Pass: false,
	}
	err := decoder.Decode(&test)
	log.Println(test)
	if err != nil {
		// TODO: Need to remove this and properly handle the error
		log.Println("couldn't decode request body:", err)
		return
	}

	if test.Pass {
		newUserLevel := models.UserLevel{test.Level, time.Now()}

		err := ctx.User.UpdateLevel(newUserLevel)
		if err != nil {
			log.Println("couldn't update user info:", err)
		}

	}

}
