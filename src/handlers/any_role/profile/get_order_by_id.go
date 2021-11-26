package profile_handler

import (
	"comiditapp/api/database"
	"comiditapp/api/env"
	"comiditapp/api/models"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderById(c *gin.Context) {
	db := database.GetDB()
	col := db.Client.Database(env.DB_NAME).Collection("orders")

	orderId := c.Param("id")
	orderObjectID, err := primitive.ObjectIDFromHex(orderId)

	if err != nil {
		panic(err)
	}

	var result models.Order
	findErr := col.FindOne(context.TODO(), bson.M{"id": orderObjectID}).Decode(&result)

	if findErr != nil {
		fmt.Println("Error calling FindOne() on GetOrderById:", findErr)
		panic(findErr)
	}

	c.JSON(http.StatusOK, result)
}
