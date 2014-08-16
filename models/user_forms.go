package models

import (
	"net/http"
	"strings"

	"log"

	"bitbucket.com/abijr/kails/util"
	"github.com/martini-contrib/binding"
	"labix.org/v2/mgo/bson"
)

/* User forms */

// User signup form
type UserSignupForm struct {
	Username string `bson:"name" form:"username" binding:"required"`
	Email    string `bson:"email" form:"email" binding:"required"`
	Password string `bson:"pass" form:"password" binding:"required"`
}

// User login form
type UserLoginForm struct {
	Email    string `bson:"email" form:"email" binding:"required"`
	Password string `bson:"pass" form:"password" binding:"required"`
}

// Validation errors
var (
	errUserAlreadyExist = binding.Error{
		FieldNames:     []string{"username"},
		Classification: "DuplicateError",
		Message:        "Username is already in use",
	}
	errEmailAlreadyExist = binding.Error{
		FieldNames:     []string{"email"},
		Classification: "DuplicateError",
		Message:        "Email is already in use",
	}
	errEmailInvalid = binding.Error{
		FieldNames:     []string{"email"},
		Classification: "InvalidError",
		Message:        "Email is invalid",
	}
	errPasswordTooShort = binding.Error{
		FieldNames:     []string{"password"},
		Classification: "PasswordError",
		Message:        "Password too short, must be at least 8 characters long",
	}

	errWrongEmailOrPassword = binding.Error{
		FieldNames:     []string{"password", "username"},
		Classification: "AuthError",
		Message:        "Wrong username or password",
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
			errors = append(errors, errUserAlreadyExist)
		}

		if uf.Email == result.Email {
			errors = append(errors, errEmailAlreadyExist)
		}
	}

	if len(uf.Password) < 8 {
		errors = append(errors, errPasswordTooShort)
	}
	return errors
}

// Validate checks if user exists and if password correct returns same error for all errors
func (ulif UserLoginForm) Validate(errors binding.Errors, req *http.Request) binding.Errors {
	var user User

	// Get user password
	err := users.
		Find(bson.M{"email": ulif.Email}).
		Select(bson.M{"pass": 1, "salt": 1}).
		One(&user)

	if err != nil {
		log.Println(err)
		errors = append(errors, errWrongEmailOrPassword)
		return errors
	}

	inputPassword := util.HashPassword(ulif.Password, user.Salt)

	if cmp := util.PasswordCompare(inputPassword, user.Password); cmp != 1 {
		log.Println("wrong password")
		errors = append(errors, errWrongEmailOrPassword)
	}

	return errors
}
