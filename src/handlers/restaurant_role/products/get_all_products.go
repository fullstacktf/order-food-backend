package product_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	c.String(http.StatusOK, "GetAll handler")
}
