package client

import (
	repository "comiditapp/api/database/repositories/users"
	client_handler "comiditapp/api/handlers/any_role/clients"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	var usersRepository repository.MockedUsersRepository

	r.GET("/clients", client_handler.FindClients(usersRepository))
}
