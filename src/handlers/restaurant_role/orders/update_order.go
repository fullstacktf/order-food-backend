package order_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOrder(c *gin.Context) {
	c.String(http.StatusOK, "UpdateOrder handler")
}
