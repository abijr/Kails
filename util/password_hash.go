package util

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"

	"code.google.com/p/go.crypto/pbkdf2"
	"crypto/subtle"
)

var (
	ErrSaltNotGenerated  = errors.New("Salt could not be generated")
	ErrPasswordNotHashed = errors.New("Password could not be hashed")
)

const (
	iterations = 4096
	saltSize   = 10
)

func HashPassword(password string, salt []byte) []byte {
	return pbkdf2.Key([]byte(password), salt, iterations, sha256.Size, sha256.New)
}

func NewSalt() ([]byte, error) {

	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, ErrSaltNotGenerated
	}

	return salt, nil
}

var PasswordCompare = subtle.ConstantTimeCompare
