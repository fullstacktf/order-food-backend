package profile_handler

import (
	repository "comiditapp/api/database/repositories/orders"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderById(repository repository.MockedOrdersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, repository.GetOrderById(context))
	}
}
