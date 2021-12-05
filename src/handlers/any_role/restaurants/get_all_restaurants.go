package restaurant_handler

import (
	repository "comiditapp/api/database/repositories/users"

	"github.com/gin-gonic/gin"
)

func FindRestaurants(repository *repository.MongoUsersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		statusCode, response := repository.FindRestaurants(context)
		context.IndentedJSON(statusCode, response)
	}
}
