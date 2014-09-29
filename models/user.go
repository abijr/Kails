package models

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"encoding/base64"

	"bitbucket.com/abijr/kails/db"
	"bitbucket.com/abijr/kails/util"
	"github.com/diegogub/aranGO"
	"gopkg.in/mgo.v2/bson"
)

const (
	// UserCollection is the name of the database collection holding user information
	UserCollection = "users"
)

// UserLevel is the representation of the
// user progress in a level
type UserLevel struct {
	// Id is the level Id
	Id int `json:"Id"`
	// LastPracticed is the list practiced time
	LastPracticed time.Time `json"Last"`
}

// User is the user structure, it holds user information
type User struct {
	aranGO.Document
	Username          string               `json:"Username"`
	Email             string               `json:"Email"`
	Password          string               `json:"Password"`
	Salt              string               `json:"Salt"`
	InterfaceLanguage string               `json:"InterfaceLanguage"`
	StudyLanguage     string               `json:"StudyLanguage"`
	Since             time.Time            `json:"Since"`
	Levels            map[string]UserLevel `json:"Levels"`
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
	hash := util.HashPassword(uf.Password, salt)
	user := User{}

	// TODO: add study language and webpage language
	// settings here
	user.Username = uf.Username
	user.Email = uf.Email
	user.Password = base64.StdEncoding.EncodeToString(hash)
	user.Salt = base64.StdEncoding.EncodeToString(salt)
	user.Since = time.Now()

	err := users.Save(user)
	if err != nil {
		return err
	}
	return nil
}

// UserSearch does a fulltext prefix search in the database
// with the argument as search string populates a slice and returns it.
func UserSearch(name string) ([]User, error) {

	var results []User
	// results := make([]User, 5)
	// If empty name, return error
	if name == "" {
		return nil, errUserNotExist
	}

	// TODO: remove bson dependency
	cur, err := users.FullText("prefix:"+name, "Username", 0, 5)
	if err != nil {
		return nil, err
	}
	cur.FetchBatch(&results)
	if cur.Count() == 0 {
		return nil, errors.New("no users")
	}
	return results, nil
}

func UserByKey(key string) (*User, error) {
	user := new(User)

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
func UserByName(name string) (*User, error) {
	user := new(User)

	// If empty name, return error
	if name == "" {
		return nil, errUserNotExist
	}

	err := users.First(bson.M{"Username": name}, user)
	if err != nil {
		return nil, errUserNotExist
	}
	return user, nil
}

// UserByName searches in the database for the user 'email' and
// populates User struct, than returns a pointer.
func UserByEmail(email string) (*User, error) {
	user := new(User)

	// If empty name, return error
	if email == "" {
		return nil, errUserNotExist
	}

	err := users.First(bson.M{"Email": email}, user)
	if err != nil {
		return nil, errUserNotExist
	}
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
