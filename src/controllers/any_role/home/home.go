package home

import (
	"comiditapp/api/database"
	home_handler "comiditapp/api/handlers/any_role/home"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	// Not sure de si la inyeccion de la db aqui tiene sentido, a futuro lo veremos
	homeGroup := r.Group("/")
	{
		homeGroup.GET("", home_handler.GetHome)
		homeGroup.GET("/home", home_handler.GetHome)
	}
}
