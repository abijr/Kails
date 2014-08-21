package models

import (
	"errors"
	"regexp"
	"time"

	"log"

	"labix.org/v2/mgo/bson"

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

// User is the user structure, it holds user information
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

// NewUser creates a new user using the UserSignupForm information
// and stores it in the database.
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

// UserByName searches in the database for the user 'name' and
// populates User struct, than returns a pointer.
func UserByName(name string) (*User, error) {
	var user *User
	user = new(User)

	// If empty name, return error
	if name == "" {
		return nil, errUserNotExist
	}

	err := users.Find(bson.M{"name": name}).One(user)
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

	err := users.Find(bson.M{"email": email}).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
