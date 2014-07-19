package models

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/martini-contrib/binding"

	"labix.org/v2/mgo/bson"

	"bitbucket.com/abijr/kails/db"
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
	Password string    "pass"
	Language string    "lang"
	Created  time.Time "since"
	Levels   []Level   "levels"
}

type UserForm struct {
	Username string `bson:"name" form:"username" binding:"required"`
	Email    string `bson:"email" form:"email" binding:"required"`
	Password string `bson:"pass" form:"password" binding:"required"`
}

// Utility variables
var (
	// users collection
	users = db.Collection(UserCollection)

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
	ErrPasswordTooShort = binding.Error{
		FieldNames:     []string{"password"},
		Classification: "DuplicateError",
		Message:        "Password too short, must be at least 8 characters long",
	}
)

func (uf UserForm) Validate(errors binding.Errors, req *http.Request) binding.Errors {

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
