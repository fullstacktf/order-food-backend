package repository

import (
	"comiditapp/api/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var orders []models.User = []models.User{
	{Id: primitive.NewObjectID(), Role: "client", Name: "Francisco Díaz", Email: "franDiaz@yahoo.es", HashedPassword: "MichaelHighFives", Phone: 111111111, Address: []models.Address{{Street: "Calle Roquefort 32", ZipCode: 11111, Country: "España", City: "Las Palmas de Gran Canaria"}}},
	{Id: primitive.NewObjectID(), Role: "client", Name: "María Fernández", Email: "mery-mail@hotmail.com", HashedPassword: "BrooklynDisguises", Phone: 222222222, Address: []models.Address{{Street: "Calle Benito 41", ZipCode: 22222, Country: "España", City: "Córdoba"}}},
	{Id: primitive.NewObjectID(), Role: "restaurant", Name: "Cafetería Mortadelo", Email: "cafeteriamortadelo@hotmail.es", HashedPassword: "HappyChickens", Phone: 111111111, Address: []models.Address{{Street: "Avenida de los Majuelos 117", ZipCode: 11111, Country: "España", City: "Santa Cruz de Tenerife"}},
		Menu: []models.Product{{Id: "1", Category: "mainCourse", Name: "Bocadillo Clásico Embutidos", Price: 1.40}, {Id: "2", Category: "mainCourse", Name: "Perrito Especial Mexicano", Price: 3.00}, {Id: "3", Category: "appetizer", Name: "Ensaladilla", Price: 4.20}}},
	{Id: primitive.NewObjectID(), Role: "restaurant", Name: "Don Perrito", Email: "DonPerrito@hotmail.com", HashedPassword: "r,<C:RW(x+2xW}{~", Phone: 222222222, Address: []models.Address{{Street: "Calle Decano Consular Jesus Ramos González 18", ZipCode: 22222, Country: "España", City: "Santa Cruz de Tenerife"}},
		Menu: []models.Product{{Id: "1", Category: "mainCourse", Name: "Ensalada Don Perrito", Price: 4.00}, {Id: "2", Category: "mainCourse", Name: "Ensalada Vegetal", Price: 4.20}, {Id: "3", Category: "drink", Name: "Cerveza Lata", Price: 1.60}}},
}

type MockedUsersRepository struct {
	mock.Mock
}

// any role
func (r *MockedUsersRepository) SignUpUser(context *gin.Context) models.User {

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) SignInUser(context *gin.Context) models.User {

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) UpdateProfile(context *gin.Context) models.User {

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) FindRestaurants(context *gin.Context) models.User {

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) GetRestaurantById(context *gin.Context) models.User {

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) GetRestaurantProducts(context *gin.Context) []models.Product {

	args := r.Called(context)
	return args.Get(0).([]models.Product)
}

// restaurant role

func (r *MockedUsersRepository) FindProducts(context *gin.Context) models.User {

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) CreateProduct(context *gin.Context) models.User {

	args := r.Called(context)
	return args.Get(0).(models.User)
}

func (r *MockedUsersRepository) UpdateProduct(context *gin.Context) models.User {

	args := r.Called(context)
	return args.Get(0).(models.User)
}
