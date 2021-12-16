package controllers

import (
	handlers "comiditapp/api/handlers/home"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	homeGroup := r.Group("/")
	{
		homeGroup.GET("", handlers.GetHome)
		homeGroup.GET("/home", handlers.GetHome)
	}
}
