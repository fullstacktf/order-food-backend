package services

import "github.com/dgrijalva/jwt-go"

// TODO: try to find a way to retrieve payload, or just return the id, that is the one we're using
func GetTokenPayload(token string) jwt.Claims {
	t, _ := jwt.Parse(token, nil)
	return t.Claims.(jwt.MapClaims)
}
