package order_handler

import (
	repository "comiditapp/api/database/repositories/orders"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetClientsOrders(repository repository.MockedOrdersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, repository.GetClientsOrders(context))
	}
}
