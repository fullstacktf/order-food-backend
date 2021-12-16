package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to comiditapp homepage!! 🍔")
}
