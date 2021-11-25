package restaurant_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRestaurantProducts(c *gin.Context) {
	c.String(http.StatusOK, "GetRestaurantProducts handler")
}
