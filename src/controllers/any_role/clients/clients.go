package client

import (
	repository "comiditapp/api/database/repositories/users"
	client_handler "comiditapp/api/handlers/any_role/clients"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	var usersRepository repository.MockedUsersRepository

	clientsGroup := r.Group("/clients")
	{
		clientsGroup.GET("", client_handler.FindClients(usersRepository))
		clientsGroup.GET("/:id", client_handler.GetClientById(usersRepository))
	}
}
