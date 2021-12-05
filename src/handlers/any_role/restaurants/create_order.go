package restaurant_handler

import (
	repository "comiditapp/api/database/repositories/orders"

	"github.com/gin-gonic/gin"
)

// request example: (POST) http://localhost:3000/restaurants/551137c2f9e1fac808a5f572/orders

func CreateOrder(repository *repository.MongoOrdersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		statusCode, response := repository.CreateOrder(context)
		context.IndentedJSON(statusCode, response)
	}
}
