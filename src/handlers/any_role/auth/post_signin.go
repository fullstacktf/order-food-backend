package auth_handler

import (
	repository "comiditapp/api/database/repositories/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignInUser(repository repository.MockedUsersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, repository.SignInUser(context))
	}
}