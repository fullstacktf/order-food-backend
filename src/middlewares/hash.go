package middlewares

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func VerifyPassword(userPass string, providedPass string) (err error) {
	if err := bcrypt.CompareHashAndPassword([]byte(providedPass), []byte(userPass)); err != nil {
		return err
	}

	return nil
}
