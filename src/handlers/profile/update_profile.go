package handlers

import (
	repository "comiditapp/api/database/repositories/users"

	"github.com/gin-gonic/gin"
)

func UpdateProfile(repository *repository.MongoUsersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		statusCode, response := repository.UpdateProfile(context)
		context.IndentedJSON(statusCode, response)
	}
}
