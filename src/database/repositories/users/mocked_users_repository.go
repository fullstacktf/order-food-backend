package repository

import (
	"comiditapp/api/enums"
	"comiditapp/api/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// This is just to ensure that the always have the same id's and dont generate new one after a server reload
var id1 = primitive.NewObjectID()
var id2 = primitive.NewObjectID()
var id3 = primitive.NewObjectID()
var id4 = primitive.NewObjectID()

var users []models.User = []models.User{
	{Id: id1, Role: enums.Client, Name: "Francisco Díaz", Email: "franDiaz@yahoo.es", HashedPassword: "MichaelHighFives", Phone: 111111111, Address: []models.Address{{Street: "Calle Roquefort 32", ZipCode: 11111, Country: "España", City: "Las Palmas de Gran Canaria"}}},
	{Id: id2, Role: enums.Client, Name: "María Fernández", Email: "mery-mail@hotmail.com", HashedPassword: "BrooklynDisguises", Phone: 222222222, Address: []models.Address{{Street: "Calle Benito 41", ZipCode: 22222, Country: "España", City: "Córdoba"}}},
	{Id: id3, Role: enums.Restaurant, Name: "Cafetería Mortadelo", Email: "cafeteriamortadelo@hotmail.es", HashedPassword: "HappyChickens", Phone: 111111111, Address: []models.Address{{Street: "Avenida de los Majuelos 117", ZipCode: 11111, Country: "España", City: "Santa Cruz de Tenerife"}},
		Menu: []models.Product{{Id: primitive.NewObjectID(), Category: "mainCourse", Name: "Bocadillo Clásico Embutidos", Price: 1.40}, {Id: primitive.NewObjectID(), Category: "mainCourse", Name: "Perrito Especial Mexicano", Price: 3.00}, {Id: primitive.NewObjectID(), Category: "appetizer", Name: "Ensaladilla", Price: 4.20}}},
	{Id: id4, Role: enums.Restaurant, Name: "Don Perrito", Email: "DonPerrito@hotmail.com", HashedPassword: "r,<C:RW(x+2xW}{~", Phone: 222222222, Address: []models.Address{{Street: "Calle Decano Consular Jesus Ramos González 18", ZipCode: 22222, Country: "España", City: "Santa Cruz de Tenerife"}},
		Menu: []models.Product{{Id: primitive.NewObjectID(), Category: "mainCourse", Name: "Ensalada Don Perrito", Price: 4.00}, {Id: primitive.NewObjectID(), Category: "mainCourse", Name: "Ensalada Vegetal", Price: 4.20}, {Id: primitive.NewObjectID(), Category: "drink", Name: "Cerveza Lata", Price: 1.60}}},
}

type MockedUsersRepository struct {
	mock.Mock
}

// any role
func (r *MockedUsersRepository) SignUpUser(context *gin.Context) models.User {

	var validate *validator.Validate = validator.New()
	var newUser models.User

	json.NewDecoder(context.Request.Body).Decode(&newUser)

	err := validate.Struct(newUser)
	if err != nil {
		errorMsg := "Cannot create user, required fields not provided...\n" + err.(validator.ValidationErrors).Error()

		http.Error(context.Writer, errorMsg, http.StatusBadRequest)
		panic(errorMsg)
	}

	users = append(users, newUser)
	context.IndentedJSON(http.StatusOK, newUser)

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) SignInUser(context *gin.Context) models.User {
	// needs session info

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) UpdateProfile(context *gin.Context) models.User {

	var validate *validator.Validate = validator.New()
	var newUserProfile models.User

	// this line should get the id from user session
	userId, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMsg := "Bad Request, " + err.Error()

		http.Error(context.Writer, errorMsg, http.StatusBadRequest)
		panic(err)
	}

	foundIndex := -1
	for index, value := range users {
		if value.Id == userId {
			foundIndex = index
		}
	}

	if foundIndex == -1 {
		errorMsg := "Cannot found user with ID " + userId.Hex()
		http.Error(context.Writer, errorMsg, http.StatusNotFound)
		panic(errorMsg)
	}

	json.NewDecoder(context.Request.Body).Decode(&newUserProfile)

	valErr := validate.Struct(newUserProfile)
	if valErr != nil {
		errorMsg := "Cannot update user, required fields not provided...\n" + valErr.(validator.ValidationErrors).Error()

		http.Error(context.Writer, errorMsg, http.StatusBadRequest)
		panic(valErr)
	}

	users[foundIndex] = newUserProfile

	context.IndentedJSON(http.StatusOK, newUserProfile)

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) FindRestaurants(context *gin.Context) []models.User {

	result := []models.User{}

	for _, val := range users {
		if val.Role == enums.Restaurant {
			result = append(result, val)
		}
	}

	context.IndentedJSON(http.StatusOK, result)

	args := r.Called(context)
	return args.Get(0).([]models.User)
}

func (r *MockedUsersRepository) FindClients(context *gin.Context) []models.User {

	result := []models.User{}

	for _, val := range users {
		if val.Role == enums.Client {
			result = append(result, val)
		}
	}

	context.IndentedJSON(http.StatusOK, result)

	args := r.Called(context)
	return args.Get(0).([]models.User)
}

func (r *MockedUsersRepository) GetRestaurantById(context *gin.Context) models.User {

	result := models.User{}
	restaurantId, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMsg := "Bad Request, " + err.Error()

		http.Error(context.Writer, errorMsg, http.StatusBadRequest)
		panic(err)
	}

	for _, val := range users {
		if val.Id == restaurantId {
			if val.Role == enums.Restaurant {
				result = val
				break
			} else {
				errorMsg := "Cannot found restaurant with ID " + restaurantId.String()
				http.Error(context.Writer, errorMsg, http.StatusNotFound)
				panic(errorMsg)
			}
		}
	}

	context.IndentedJSON(http.StatusOK, result)

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) GetClientById(context *gin.Context) models.User {

	result := models.User{}
	clientId, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMsg := "Bad Request, " + err.Error()

		http.Error(context.Writer, errorMsg, http.StatusBadRequest)
		panic(err)
	}

	for _, val := range users {
		if val.Id == clientId {
			if val.Role == enums.Client {
				result = val
				break
			} else {
				errorMsg := "Cannot found client with ID " + clientId.String()
				http.Error(context.Writer, errorMsg, http.StatusNotFound)
				panic(errorMsg)
			}
		}
	}

	context.IndentedJSON(http.StatusOK, result)

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) GetRestaurantProducts(context *gin.Context) []models.Product {

	result := []models.Product{}

	restaurantId, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMsg := "Bad Request, " + err.Error()

		http.Error(context.Writer, errorMsg, http.StatusBadRequest)
		panic(err)
	}

	for _, val := range users {
		if val.Id == restaurantId {
			if val.Role == enums.Restaurant {
				result = append(result, val.Menu...)
				break
			} else {
				errorMsg := "Cannot found restaurant with ID " + restaurantId.String()
				http.Error(context.Writer, errorMsg, http.StatusNotFound)
				panic(errorMsg)
			}
		}
	}

	if len(result) == 0 {
		errorMsg := "Cannot found products for restaurant with ID " + restaurantId.String()
		http.Error(context.Writer, errorMsg, http.StatusNotFound)
		panic(errorMsg)
	}

	context.IndentedJSON(http.StatusOK, result)

	args := r.Called(context)
	return args.Get(0).([]models.Product)
}

// restaurant role

func (r *MockedUsersRepository) FindProducts(context *gin.Context) models.User {
	// needs session info

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) CreateProduct(context *gin.Context) models.User {
	// needs session info

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) UpdateProduct(context *gin.Context) models.User {
	// needs session info

	args := r.Called(context)
	return args.Get(0).(models.User)
}
