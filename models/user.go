package models

import (
	"errors"
	"time"

	"bitbucket.com/abijr/kails/db"
)

const (
	UserCollection = "users"
)

type User struct {
	Name     string    "name"
	Email    string    "email"
	Password string    "pass"
	Language string    "lang"
	Created  time.Time "since"
}

var (
	// users collection
	users = db.Collection(UserCollection)

	// Errors
	ErrUserAlreadyExist = errors.New("User already exist")
	ErrUserNotExist     = errors.New("User does not exist")
	ErrEmailAlreadyUsed = errors.New("E-mail already used")
	ErrUserNameIllegal  = errors.New("User name contains illegal characters")
)

func UserByName(name string) (*User, error) {
	var user *User

	// If empty name, return error
	if name != "" {
		user = new(User)
		user.Name = name
	} else {
		return nil, ErrUserNotExist
	}

	users.Find(user.Name).One(user)
	return user, nil
}
