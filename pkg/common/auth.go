package common

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func CompareHashPassword(inputPassword string, dbPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(inputPassword))
	return err == nil
}
