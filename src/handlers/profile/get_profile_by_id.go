package handlers

import (
	repository "comiditapp/api/database/repositories/users"

	"github.com/gin-gonic/gin"
)

func GetProfileById(repository *repository.MongoUsersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		statusCode, response := repository.GetProfileById(context)
		context.IndentedJSON(statusCode, response)
	}
}
