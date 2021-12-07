package profile_handler

import (
	repository "comiditapp/api/database/repositories/orders"

	"github.com/gin-gonic/gin"
)

func FindOrders(repository *repository.MongoOrdersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		statusCode, response := repository.FindOrders(context)
		context.IndentedJSON(statusCode, response)
	}
}
