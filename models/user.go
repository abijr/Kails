package models

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"encoding/base64"
	"log"

	"bitbucket.com/abijr/kails/db"
	"bitbucket.com/abijr/kails/util"
	"github.com/diegogub/aranGO"
	"gopkg.in/mgo.v2/bson"
)

const (
	// Name of the database collection holding user information
	UserCollection = "users"
)

type UserLevel struct {
	Id            int       `json:"id"`
	LastPracticed time.Time `json"last"`
}

// User is the user structure, it holds user information
type User struct {
	aranGO.Document
	Username      string               `json:"Username"`
	Email         string               `json:"Email"`
	Password      string               `json:"Password"`
	Salt          string               `json:"Salt"`
	Language      string               `json:"Language"`
	StudyLanguage string               `json:"StudyLanguage"`
	Created       time.Time            `json:"Created"`
	Levels        map[string]UserLevel `json:"Levels"`
}

// Utility variables
var (
	// users collection
	users = db.Collection(UserCollection)

	// email regexp
	emailPattern = regexp.MustCompile("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$")

	errUserNotExist    = errors.New("user does not exist")
	errUserNameIllegal = errors.New("user name contains illegal characters")
)

// NewUser creates a new user using the UserSignupForm information
// and stores it in the database.
func NewUser(uf UserSignupForm) error {
	salt, _ := util.NewSalt()
	t0 := time.Now()
	hash := util.HashPassword(uf.Password, salt)
	log.Println("Hash time: ", time.Since(t0).String())
	user := User{}

	// TODO: add study language and webpage language
	// settings here
	user.Username = uf.Username
	user.Email = uf.Email
	user.Password = base64.StdEncoding.EncodeToString(hash)
	user.Salt = base64.StdEncoding.EncodeToString(salt)
	user.Created = time.Now()

	err := users.Save(user)
	if err != nil {
		return err
	}
	return nil
}

// UserByName searches in the database for the user 'name' and
// populates User struct, than returns a pointer.
func UserByName(name string) (*User, error) {
	var user *User
	user = new(User)

	// If empty name, return error
	if name == "" {
		return nil, errUserNotExist
	}

	// TODO: remove bson dependency
	err := users.First(bson.M{"Name": name}, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserByKey(key string) (*User, error) {
	var user *User
	user = new(User)

	if key == "" {
		return nil, errUserNotExist
	}

	err := users.Get(key, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UserByName searches in the database for the user 'email' and
// populates User struct, than returns a pointer.
func UserByEmail(email string) (*User, error) {
	var user *User
	user = new(User)

	// If empty name, return error
	if email == "" {
		return nil, errUserNotExist
	}

	cur, _ := users.Example(bson.M{"Email": email}, 0, 1)
	cur.FetchOne(user)
	return user, nil
}

func (user *User) UpdateLevel(level UserLevel) error {

	// The field to update is in the format
	// `level.{{level_number}}`
	levelString := fmt.Sprintf("%v", level.Id)

	//  query is formated as follows:
	// {$set: {"level.1": {
	// 			"last": ISODate("blah blah")
	// 		}
	// }}
	updateQuery :=
		bson.M{
			"level": bson.M{
				levelString: bson.M{
					"last": level.LastPracticed,
				},
			},
		}
	err := users.Patch(user.Key, updateQuery)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) UpdateStudyLanguage(lang string) error {

	updateQuery := bson.M{
		"StudyLanguage": lang,
	}
	err := users.Patch(user.Key, updateQuery)
	if err != nil {
		return err
	}

	return nil

}

func Test() {
	user := new(User)

	dbase, _ := aranGO.Connect("http://localhost:8529", "", "", true)
	b := dbase.DB("kails").Col("users")
	// Get user password
	cur, _ := b.Example(bson.M{"Email": "user@email.com"}, 0, 1)
	cur.FetchOne(user)

	log.Println(user)
}
