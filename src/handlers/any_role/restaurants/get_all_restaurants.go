package restaurant_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllRestaurants(c *gin.Context) {
	c.String(http.StatusOK, "GetAllRestaurants handler")
}
