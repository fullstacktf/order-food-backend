package controllers

import (
	"comiditapp/api/database"
	handlers "comiditapp/api/handlers/clients"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	clientsGroup := r.Group("/clients")
	{
		clientsGroup.GET("", handlers.FindClients(db.UsersRepository))
		clientsGroup.GET("/:id", handlers.GetClientById(db.UsersRepository))
	}
}
