package product_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostProduct(c *gin.Context) {
	c.String(http.StatusOK, "PostProduct handler")
}
