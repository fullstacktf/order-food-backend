package handlers

import (
	repository "comiditapp/api/database/repositories/users"
	"comiditapp/api/models"
	"comiditapp/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func Register(repository *repository.MongoUsersRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		var user models.User
		context.BindJSON(&user)

		validate := validator.New()

		if err := validate.Struct(user); err != nil {
			validatorError := err.(validator.ValidationErrors).Error()
			errorMessage := "Cannot create user, required fields not provided\n" + validatorError
			services.SendResponse(context, services.Response{Status: http.StatusBadRequest, Error: []string{errorMessage}})
		}

		// check if user exists
		if repository.DoesUserExists(user) == true {
			services.SendResponse(context, services.Response{Status: http.StatusBadRequest, Error: []string{"That email is already registered"}})
		} else {
			var parsedMenu []models.Product = services.CreateMenu(user.Menu)
			user.Menu = parsedMenu

			newUser, err := services.CreateUser(user)
			if err != nil {
				services.SendResponse(context, services.Response{Status: http.StatusInternalServerError, Error: []string{"Internal error on register"}})
			}

			if err := repository.Register(*newUser); err != nil {
				services.SendResponse(context, services.Response{Status: http.StatusInternalServerError, Error: []string{"Internal error on register"}})
			}

			token, err := services.SetUserCookie(context, *newUser)
			if err != nil {
				services.SendResponse(context, services.Response{Status: http.StatusInternalServerError, Error: []string{"Internal error on register"}})
			}

			services.SendResponse(context, services.Response{Status: http.StatusCreated, Data: map[string]interface{}{"user": services.UserToDomain((*newUser)), "token": token}})
		}
	}
}
