package routes

import (
	"log"
	"time"

	"github.com/go-martini/martini"

	"bitbucket.com/abijr/kails/middleware"
	"bitbucket.com/abijr/kails/models"
)

var (
	usersConnected []string
	isJustConnected bool
) 

func Home(ctx *middleware.Context) {
	if ctx.IsLogged {
		ctx.Data["Title"] = "Home"
		ctx.Data["Name"] = ctx.User.Username
		ctx.Data["Language"] = ctx.User.StudyLanguage
		ctx.Data["Country"] = "Mexico"
		//ctx.Data["Email"] = ctx.User.Email
		//ctx.Data["Since"] = ctx.User.Since.Format("January 2 of 2006")
		ctx.HTML(200, "user/home")
	} else {
		ctx.Data["Title"] = "Welcome"
		ctx.HTML(200, "main/main")
	}
}

func UserPage(ctx *middleware.Context, params martini.Params) {
	username := params["name"]
	user, err := models.UserByName(username)
	if err != nil {
		ctx.HTML(500, "")
		return
	}

	ctx.Data["Title"] = username
	ctx.Data["Username"] = user.Username
	ctx.Data["Email"] = user.Email
	ctx.Data["Since"] = user.Since.Format("January 2 of 2006")
	ctx.HTML(200, "user/info")
}

type SearchResult struct {
	Data  []string `json:"Data"`
	Error string   `json:"Error"`
}

func UserSearch(ctx *middleware.Context, params martini.Params) {
	var data SearchResult

	searchString := params["name"]
	results, err := models.UserSearch(searchString)

	if err != nil {
		data.Error = "couldn't find any users (with error)"
	}

	userList := make([]string, 0, 5)

	for _, user := range results {
		log.Println(user.Username)
		userList = append(userList, user.Username)
	}

	if len(userList) == 0 {
		data.Error = "couldn't find any users"
	} else {
		data.Data = userList
		data.Error = ""
	}

	ctx.JSON(200, data)

}

func Settings(ctx *middleware.Context) {
	ctx.Data["Title"] = "Settings"
	ctx.Data["currentLanguage"] = ctx.User.StudyLanguage
	var lang string
	if ctx.User.StudyLanguage == "english" {
		lang = "spanish"
	} else {
		lang = "english"
	}
	ctx.Data["otherLanguage"] = lang
	ctx.HTML(200, "user/settings")
}

type SettingsForm struct {
	Language string `form:"language"`
}

func SettingsPost(ctx *middleware.Context, form SettingsForm) {
	if form.Language == ctx.User.StudyLanguage {
		ctx.Redirect("/")
		return
	}

	err := ctx.User.UpdateStudyLanguage(form.Language)
	if err != nil {
		//TODO: put something here
	}

}

func SignUp(ctx *middleware.Context) {
	ctx.Data["Title"] = "Sign Up"
	ctx.HTML(200, "user/signup")
}

func SignUpPost(ctx *middleware.Context, form models.UserSignupForm) {
	err := models.NewUser(form)
	if err != nil {
		log.Println(err)
		ctx.HTML(501, "")
		return
	}
	ctx.Redirect("/")
}

func Login(ctx *middleware.Context) {
	if ctx.IsLogged {
		ctx.Redirect("/")
	} else {
		ctx.Data["Title"] = "Login"
		ctx.HTML(200, "user/login")
	}
}

func LoginPost(ctx *middleware.Context, form models.UserLoginForm) {
	ctx.Data["Title"] = "Home"
	user, err := models.UserByEmail(form.Email)
	if err != nil {
		// TODO: fill this up.
		log.Println(err)
		log.Println(user)
	}

	ctx.User = user
	ctx.Session.Set("key", user.Key)
	ctx.IsLogged = true

	usersConnected = append(usersConnected, ctx.User.Username)
	isJustConnected = true

	ctx.Redirect("/")
}

func Logout(ctx *middleware.Context) {
	if ctx.IsLogged {
		// This is necessary
		ctx.Session.Clear()

		// This is for making sure
		ctx.User = new(models.User) // blank user
		ctx.IsLogged = false

		log.Println("user logged out")
	}

	ctx.Redirect("/")
}

func Friends(ctx *middleware.Context) {
	ctx.Data["Title"] = "Friends"
	ctx.Data["Name"] = ctx.User.Username
	ctx.Data["Language"] = ctx.User.StudyLanguage
	ctx.Data["Country"] = "Mexico"
	ctx.HTML(200, "user/friends")
}

func GetFriends(ctx *middleware.Context) {
	friends, err := ctx.User.ListFriends()

	if err != nil {
		log.Println("Error: ", err)
	}

	log.Println(friends)
	ctx.JSON(200, friends)
}

func CheckFriendStatus(ctx *middleware.Context) {
	friend := make(chan string, 1)

	lenght := len(usersConnected)

	go func() {
		log.Println(isJustConnected)
		if isJustConnected {
			friend <- usersConnected[lenght- 1]
		}
	}()

	select {
		case res := <-friend:
			isJustConnected = false
			if res == "other" {
				//ctx.JSON(200, res)
				log.Println("################################################")
				log.Println("Found: ", res)
				log.Println("################################################")
			}
		case <-time.After(time.Second * 60):
			//ctx.JSON(200, friend)
			log.Println("################################################")
			log.Println("Not Found")
			log.Println("################################################")
	}
}