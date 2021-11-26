package auth_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSignIn(c *gin.Context) {
	c.String(http.StatusOK, "PostSignIn handler")
}
