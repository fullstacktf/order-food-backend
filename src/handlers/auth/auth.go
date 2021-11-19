package auth

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	auth := r.Group("auth")
	auth.POST("/auth/signup/restaurant", PostSignUpRestaurant)
	auth.POST("/auth/signin/restaurant", PostSignInRestaurant)
	auth.POST("/auth/signup/client", PostSignUpClient)
	auth.POST("/auth/signin/client", PostSignInClient)
}

func PostSignUpRestaurant(c *gin.Context) {}
func PostSignInRestaurant(c *gin.Context) {}
func PostSignUpClient(c *gin.Context)     {}
func PostSignInClient(c *gin.Context)     {}
