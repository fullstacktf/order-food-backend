package restaurant_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostOrder(c *gin.Context) {
	c.String(http.StatusOK, "PostOrder handler")
}
