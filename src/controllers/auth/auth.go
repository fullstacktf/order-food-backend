package controllers

import (
	"comiditapp/api/database"
	handlers "comiditapp/api/handlers/auth"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, db database.DB) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", handlers.Register(db.UsersRepository))
		authGroup.POST("/login", handlers.Login(db.UsersRepository))
		authGroup.POST("/logout", handlers.Logout(db.UsersRepository))
	}
}
