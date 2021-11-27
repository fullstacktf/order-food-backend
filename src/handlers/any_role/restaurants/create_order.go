package restaurant_handler

import (
	repository "comiditapp/api/database/repositories/orders"
	"net/http"

	"github.com/gin-gonic/gin"
)

// request example: (POST) http://localhost:3000/restaurants/551137c2f9e1fac808a5f572/orders

func CreateOrder(repository repository.MockedOrdersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, repository.CreateOrder(context))
	}
}
