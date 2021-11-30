package client_handler

import (
	repository "comiditapp/api/database/repositories/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetClientById(repository repository.MongoUsersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, repository.GetClientById(context))
	}
}
