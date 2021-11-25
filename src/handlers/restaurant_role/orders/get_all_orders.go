package order_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(c *gin.Context) {
	c.String(http.StatusOK, "GetAllOrders handler")
}
