package product_handler

import (
	repository "comiditapp/api/database/repositories/users"

	"github.com/gin-gonic/gin"
)

func CreateProduct(repository *repository.MongoUsersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		statusCode, response := repository.CreateProduct(context)
		context.IndentedJSON(statusCode, response)
	}
}
