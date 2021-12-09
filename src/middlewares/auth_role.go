package middlewares

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthClientIsRequired(c *gin.Context) {
	session := sessions.Default(c)
	role := session.Get("role")
	if role == "client" {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
}

func AuthRestaurantIsRequired(c *gin.Context) {
	session := sessions.Default(c)
	role := session.Get("role")
	if role == "restaurant" {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
}
