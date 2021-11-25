package profile_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderById(c *gin.Context) {
	c.String(http.StatusOK, "GetOrderById handler")
}
