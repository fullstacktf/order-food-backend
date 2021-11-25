package profile_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProfile(c *gin.Context) {
	c.String(http.StatusOK, "UpdateProfile handler")
}
