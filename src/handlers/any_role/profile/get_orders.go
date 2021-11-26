package profile_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	c.String(http.StatusOK, "GetOrders handler")
}
