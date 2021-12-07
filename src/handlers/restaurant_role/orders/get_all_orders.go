package order_handler

import (
	repository "comiditapp/api/database/repositories/orders"

	"github.com/gin-gonic/gin"
)

func FindClientOrders(repository *repository.MongoOrdersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		statusCode, response := repository.FindClientOrders(context)
		context.IndentedJSON(statusCode, response)
	}
}
