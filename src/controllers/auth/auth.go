package controllers

import (
	"comiditapp/api/database"
	handlers "comiditapp/api/handlers/auth"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", handlers.SignUpUser(db.UsersRepository))
		authGroup.POST("/signin", handlers.SignInUser(db.UsersRepository))
		authGroup.POST("/signout", handlers.SignOutUser(db.UsersRepository))
	}
}
