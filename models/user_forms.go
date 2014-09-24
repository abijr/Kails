package models

import (
	"encoding/base64"
	"net/http"
	"strings"

	"log"

	"bitbucket.com/abijr/kails/util"
	"github.com/martini-contrib/binding"
)

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

// User signup form
type UserSignupForm struct {
	Username string `form:"username" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (uf UserSignupForm) Validate(errors binding.Errors, req *http.Request) binding.Errors {

	uf.Username = strings.TrimSpace(uf.Username)
	uf.Email = strings.TrimSpace(uf.Email)

	// query := bson.M{
	// 	"$or": [2]bson.M{
	// 		bson.M{"name": uf.Username},
	// 		bson.M{"email": uf.Email},
	// 	},
	// }

	// TODO: Find which values not unique before
	// or after insertion
	// // Find if user or email already in db
	// docs := users.
	// 	Find(query).
	// 	Select(bson.M{"name": 1, "email": 1}).
	// 	Limit(2).
	// 	Iter()

	// var result User
	// for docs.Next(&result) {
	// 	if uf.Username == result.Username {
	// 		errors = append(errors, errUserAlreadyExist)
	// 	}
	//
	// 	if uf.Email == result.Email {
	// 		errors = append(errors, errEmailAlreadyExist)
	// 	}
	// }

	// TODO: put password lenght in a constant
	// Min password lenght
	if len(uf.Password) < 8 {
		errors = append(errors, errPasswordTooShort)
	}
	return errors
}

// User login form
type UserLoginForm struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// Validate checks if user exists and if password correct returns same error for all errors
func (ulif UserLoginForm) Validate(errors binding.Errors, req *http.Request) binding.Errors {
	var user *User

	// Get user password
	user, err := UserByEmail(ulif.Email)

	log.Println(user)
	if err != nil {
		log.Println(err)
		errors = append(errors, errWrongEmailOrPassword)
		return errors
	}

	// Need to convert strings to bytes
	pass, err1 := base64.StdEncoding.DecodeString(user.Password)
	salt, err2 := base64.StdEncoding.DecodeString(user.Salt)
	if err1 != nil || err2 != nil {
		log.Println("Couldn't decode strings: ")
		log.Println("pass: ", err1)
		log.Println("salt: ", err2)
	}

	inputPassword := util.HashPassword(ulif.Password, salt)

	log.Println("Input Hash: ", inputPassword, "Stored Hash: ", pass)

	if cmp := util.PasswordCompare(inputPassword, pass); cmp != 1 {
		log.Println("wrong password")
		errors = append(errors, errWrongEmailOrPassword)
	}

	return errors
}
