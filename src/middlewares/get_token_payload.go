package middlewares

import "github.com/dgrijalva/jwt-go"

func GetTokenPayload(token string) jwt.Claims {
	t, _ := jwt.Parse(token, nil)
	return t.Claims.(jwt.MapClaims)
}
