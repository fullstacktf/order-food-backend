package restaurant_handler

import (
	"comiditapp/api/database"
	"comiditapp/api/enums"
	"comiditapp/api/env"
	"comiditapp/api/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// request example: (POST) http://localhost:3000/restaurants/551137c2f9e1fac808a5f572/orders
func PostOrder(c *gin.Context) {
	db := database.GetDB()
	col := db.Client.Database(env.DB_NAME).Collection("orders")

	restaurantId := c.Param("id")

	restaurantObjectID, err := primitive.ObjectIDFromHex(restaurantId)

	if err != nil {
		panic(err)
	}

	newOrder := models.Order{
		Id:           primitive.NewObjectID(),
		ClientId:     primitive.NewObjectID(), // should get it from actual user
		RestaurantId: restaurantObjectID,
		Status:       enums.Ordered,
		TotalPrice:   21.21,
		Products: []models.ProductInfo{
			{ProductId: primitive.NewObjectID(), Quantity: 1},
		},
	}

	res, err := col.InsertOne(context.TODO(), newOrder)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, res)
	// c.String(http.StatusOK, "PostOrder handler")
}
