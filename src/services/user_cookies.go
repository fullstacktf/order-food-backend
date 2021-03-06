package services

import (
	"comiditapp/api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetUserCookie(context *gin.Context, user models.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 8760)
	token, err := GenerateJWT(
		user.Email,
		user.Id.Hex(),
		string(user.Role),
		expirationTime.Unix(),
	)
	if err != nil {
		return "", err
	}

	c := &http.Cookie{
		Name:    "token",
		Value:   token,
		Path:    "/",
		Expires: expirationTime,
	}
	http.SetCookie(context.Writer, c)

	return token, nil
}

func UnsetUserCookie(context *gin.Context) {
	c := &http.Cookie{
		Name:    "token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}

	http.SetCookie(context.Writer, c)
}
