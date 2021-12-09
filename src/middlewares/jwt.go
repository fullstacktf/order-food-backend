package middlewares

import (
	"comiditapp/api/env"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(email, id, role string, expirationTime int64) (response string, err error) {

	var jwtSecretKey = []byte(env.SECRET)

	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = expirationTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
