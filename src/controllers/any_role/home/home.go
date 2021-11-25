package home

import (
	home_handler "comiditapp/api/handlers/any_role/home"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	homeGroup := r.Group("/")
	{
		homeGroup.GET("", home_handler.GetHome)
		homeGroup.GET("/home", home_handler.GetHome)
	}
}
