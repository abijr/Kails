package models

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/martini-contrib/binding"

	"log"

	"labix.org/v2/mgo/bson"

	"bitbucket.com/abijr/kails/db"
	"bitbucket.com/abijr/kails/util"
)

const (
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

/* User forms */

// User signup form
type UserSignupForm struct {
	Username string `bson:"name" form:"username" binding:"required"`
	Email    string `bson:"email" form:"email" binding:"required"`
	Password string `bson:"pass" form:"password" binding:"required"`
}

// Utility variables
var (
	// users collection
	users = db.Collection(UserCollection)

	// email regexp
	emailPattern = regexp.MustCompile("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$")

	ErrUserNotExist    = errors.New("User does not exist")
	ErrUserNameIllegal = errors.New("User name contains illegal characters")
)

// Validation Errors
var (
	ErrUserAlreadyExist = binding.Error{
		FieldNames:     []string{"username"},
		Classification: "DuplicateError",
		Message:        "Username is already in use",
	}
	ErrEmailAlreadyExist = binding.Error{
		FieldNames:     []string{"email"},
		Classification: "DuplicateError",
		Message:        "Email is already in use",
	}
	ErrEmailInvalid = binding.Error{
		FieldNames:     []string{"email"},
		Classification: "InvalidError",
		Message:        "Email is invalid",
	}
	ErrPasswordTooShort = binding.Error{
		FieldNames:     []string{"password"},
		Classification: "PasswordError",
		Message:        "Password too short, must be at least 8 characters long",
	}
)

func (uf UserSignupForm) Validate(errors binding.Errors, req *http.Request) binding.Errors {

	uf.Username = strings.TrimSpace(uf.Username)
	uf.Email = strings.TrimSpace(uf.Email)

	query := bson.M{
		"$or": [2]bson.M{
			bson.M{"name": uf.Username},
			bson.M{"email": uf.Email},
		},
	}

	// Find if user or email already in db
	docs := users.
		Find(query).
		Select(bson.M{"name": 1, "email": 1}).
		Limit(2).
		Iter()

	var result User
	for docs.Next(&result) {
		if uf.Username == result.Username {
			errors = append(errors, ErrUserAlreadyExist)
		}

		if uf.Email == result.Email {
			errors = append(errors, ErrEmailAlreadyExist)
		}
	}

	if len(uf.Password) < 8 {
		errors = append(errors, ErrPasswordTooShort)
	}
	return errors
}

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

func AuthUser(uf UserLoginForm) {
	var user &User
	users.Find(uf.Email).One(user)
	


}

func UserByName(name string) (*User, error) {
	var user *User

	// If empty name, return error
	if name != "" {
		user = new(User)
		user.Username = name
	} else {
		return nil, ErrUserNotExist
	}

	users.Find(user.Username).One(user)
	return user, nil
}

func UserByEmail(email string) (*User, error) {
	return nil, nil
}
