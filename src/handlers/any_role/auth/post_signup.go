package auth_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSignUp(c *gin.Context) {
	c.String(http.StatusOK, "PostSignUp handler")
}
