package restaurant_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// request example: (POST) http://localhost:3000/restaurants/551137c2f9e1fac808a5f572/orders
func PostOrder(c *gin.Context) {
	c.JSON(http.StatusCreated, "res")
}
