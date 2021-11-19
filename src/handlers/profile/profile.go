package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfileOrders(c *gin.Context) {
	name := c.Param("id")
	// Logica
	message := "all good" + name
	c.String(http.StatusOK, message)
}

func GetProfileOrderById(c *gin.Context) {
	name := c.Param("id")
	// Logica
	message := "all good" + name
	c.String(http.StatusOK, message)
}

func UpdateProfileData(c *gin.Context) {}
