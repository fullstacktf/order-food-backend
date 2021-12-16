package handlers

import (
	repository "comiditapp/api/database/repositories/users"

	"github.com/gin-gonic/gin"
)

func UpdateProduct(repository *repository.MongoUsersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		statusCode, response := repository.UpdateProduct(context)
		context.IndentedJSON(statusCode, response)
	}
}
