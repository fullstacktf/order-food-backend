package product_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProduct(c *gin.Context) {
	c.String(http.StatusOK, "UpdateProduct handler")
}
