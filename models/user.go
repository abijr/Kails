package models

import (
	"errors"
	"regexp"
	"time"

	"log"

	"bitbucket.com/abijr/kails/db"
	"bitbucket.com/abijr/kails/util"
)

const (
	// Name of the database collection holding user information
	UserCollection = "users"
)

type Level struct {
	Id            int       "id"
	Name          string    "name"
	LastPracticed time.Time "last"
}

type User struct {
	Username string    "name"
	Email    string    "email"
	Password []byte    "pass"
	Salt     []byte    "salt"
	Language string    "lang"
	Created  time.Time "since"
	Levels   []Level   "levels"
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

func NewUser(uf UserSignupForm) error {
	salt, _ := util.NewSalt()
	t0 := time.Now()
	hash := util.HashPassword(uf.Password, salt)
	log.Println("Hash time: ", time.Since(t0).String())
	user := User{}

	user.Username = uf.Username
	user.Email = uf.Email
	user.Password = hash
	user.Salt = salt
	user.Created = time.Now()

	err := users.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func UserByName(name string) (*User, error) {
	var user *User

	// If empty name, return error
	if name != "" {
		user = new(User)
		user.Username = name
	} else {
		return nil, errUserNotExist
	}

	// TODO: Error check user find here
	users.Find(user.Username).One(user)
	return user, nil
}

func UserByEmail(email string) (*User, error) {
	return nil, nil
}
