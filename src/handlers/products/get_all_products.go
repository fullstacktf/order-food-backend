package handlers

import (
	repository "comiditapp/api/database/repositories/users"

	"github.com/gin-gonic/gin"
)

func FindProducts(repository *repository.MongoUsersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		statusCode, response := repository.FindProducts(context)
		context.IndentedJSON(statusCode, response)
	}
}
