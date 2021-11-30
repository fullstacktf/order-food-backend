package client

import (
	"comiditapp/api/database"
	client_handler "comiditapp/api/handlers/any_role/clients"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	clientsGroup := r.Group("/clients")
	{
		clientsGroup.GET("", client_handler.FindClients(db.UsersRepository))
		clientsGroup.GET("/:id", client_handler.GetClientById(db.UsersRepository))
	}
}
