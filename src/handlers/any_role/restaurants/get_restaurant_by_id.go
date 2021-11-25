package restaurant_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRestaurantById(c *gin.Context) {
	c.String(http.StatusOK, "GetRestaurantById handler")
}
