package repository

import (
	"comiditapp/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOrdersRepository struct {
	orders *mongo.Collection
}

func NewMongoOrdersRepository(db *mongo.Database) *MongoOrdersRepository {
	return &MongoOrdersRepository{orders: db.Collection("orders")}
}

// any_role methods

// GET - http://localhost:3000/profile/orders
// Actualmente devuelve todas las orders.
// Debería usar información de la sesión para devolver solo las del user concreto.
// Otra opcion es que se pidan todas, y luego en front se filtren las del usuario.
func (r *MongoOrdersRepository) FindOrders(context *gin.Context) (statusCode int, response interface{}) {
	foundOrders, err := r.orders.Find(context, bson.M{})
	if err != nil {
		return http.StatusConflict, err.Error()
	}

	orders := []*models.Order{}
	if err := foundOrders.All(context, &orders); err != nil {
		return http.StatusConflict, err.Error()
	}

	return http.StatusOK, orders
}

// GET - http://localhost:3000/profile/orders/:id
// Devuelve bien la order concreta.
// Debería usar información de la sesión para devolver la order solo si pertenece al user.
// Otra opcion es que se pida, y que luego en front no se le dé al user si no es de el.
func (r *MongoOrdersRepository) GetOrderById(context *gin.Context) (statusCode int, response interface{}) {
	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, errorMessage
	}

	filter := bson.M{"id": id}
	foundOrders, err := r.orders.Find(context, filter)
	if err != nil {
		return http.StatusConflict, err.Error()
	}

	orders := []*models.Order{}
	if err := foundOrders.All(context, &orders); err != nil {
		return http.StatusConflict, err.Error()
	}

	return http.StatusOK, orders
}

// Para que un usuario cree una order, deberia existir el mismo, deberia existir el restaurante sobre el que va
// a efectuar la order. La info del user que hace la order la sacariamos de la sesion
// POST - http://localhost:3000/restaurants/:id/order
// Actualmente crea una order, pero esta debería sacar el id del cliente de la sesion
// y el id del restaurante no es requerido porque lo sacamos de :id
func (r *MongoOrdersRepository) CreateOrder(context *gin.Context) (statusCode int, response interface{}) {
	var validate *validator.Validate = validator.New()
	var newOrder models.Order

	context.BindJSON(&newOrder)

	err := validate.Struct(newOrder)
	if err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot create order, required fields not provided\n" + validatorError
		return http.StatusBadRequest, errorMessage
	}

	restaurantId, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, errorMessage
	}

	id := primitive.NewObjectID()
	order := &models.Order{
		Id:           id,
		ClientId:     newOrder.ClientId,
		RestaurantId: restaurantId,
		Status:       newOrder.Status,
		TotalPrice:   newOrder.TotalPrice,
		Products:     newOrder.Products,
	}

	if _, err := r.orders.InsertOne(context, order); err != nil {
		return http.StatusBadRequest, err.Error()
	}

	return http.StatusCreated, id.Hex()
}

// restaurant_role methods

// Para esto necesitamos los permisos del restaurante para realizar las acciones, se sacará del JWT
// PUT - http://localhost:3000/orders/:id
// Actualmente actualiza bien la order, pero esta acción debería estar limitada al restaurante encargado
// de dicha order.
// Otro aspecto a tener en cuenta, es que cuando se implemente el JWT, ninguno de los ID de cliente, order, o restaurante
// deberian ser obligatorios, de manera que el usuario solo actualizará el resto de campos, ya que los de id no le incumben
func (r *MongoOrdersRepository) UpdateClientOrder(context *gin.Context) (statusCode int, response interface{}) {
	var validate *validator.Validate = validator.New()
	var newOrder models.Order

	context.BindJSON(&newOrder)

	err := validate.Struct(newOrder)
	if err != nil {
		validatorError := err.(validator.ValidationErrors).Error()
		errorMessage := "Cannot update order, required fields not provided\n" + validatorError
		return http.StatusBadRequest, errorMessage
	}

	id, err := primitive.ObjectIDFromHex(context.Param("id"))
	if err != nil {
		errorMessage := "Bad request, " + context.Param("id") + " is not a valid ID"
		return http.StatusBadRequest, errorMessage
	}

	filter := bson.M{"id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"status": newOrder.Status}}
	if _, err := r.orders.UpdateOne(context, filter, update); err != nil {
		return http.StatusBadRequest, err.Error()
	}

	return http.StatusOK, id.Hex()
}

// Para esto necesitamos los permisos del restaurante para realizar las acciones, se sacará del JWT
// GET - http://localhost:3000/orders
// Este endpoint es como el de devolver todas las orders, solo que cuando tengamos JWT sabremos el restaurante
// cuyas orders hay que traer. No lo hago porque es casi idéntico, esperaré a tener JWT
func (r *MongoOrdersRepository) FindClientOrders(context *gin.Context) (statusCode int, response interface{}) {
	return 0, &[]models.Order{}
}
