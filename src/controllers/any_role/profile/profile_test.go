package profile

import (
	"bytes"
	repository "comiditapp/api/database/repositories/orders"
	"comiditapp/api/enums"
	profile_handler "comiditapp/api/handlers/any_role/profile"
	"comiditapp/api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var repositoryMock repository.MockedOrdersRepository

var ordersMock []models.Order = []models.Order{
	{Id: primitive.NewObjectID(), RestaurantId: primitive.NewObjectID(), ClientId: primitive.NewObjectID(), Status: enums.Ordered, TotalPrice: 30.30, Products: []models.ProductInfo{{ProductId: primitive.NewObjectID(), Quantity: 1}}},
	{Id: primitive.NewObjectID(), RestaurantId: primitive.NewObjectID(), ClientId: primitive.NewObjectID(), Status: enums.Preparing, TotalPrice: 45.30, Products: []models.ProductInfo{{ProductId: primitive.NewObjectID(), Quantity: 2}, {ProductId: primitive.NewObjectID(), Quantity: 2}}},
	{Id: primitive.NewObjectID(), RestaurantId: primitive.NewObjectID(), ClientId: primitive.NewObjectID(), Status: enums.Preparing, TotalPrice: 45.30, Products: []models.ProductInfo{{ProductId: primitive.NewObjectID(), Quantity: 2}, {ProductId: primitive.NewObjectID(), Quantity: 2}}},
	{Id: primitive.NewObjectID(), RestaurantId: primitive.NewObjectID(), ClientId: primitive.NewObjectID(), Status: enums.Preparing, TotalPrice: 45.30, Products: []models.ProductInfo{{ProductId: primitive.NewObjectID(), Quantity: 2}, {ProductId: primitive.NewObjectID(), Quantity: 2}}},
}

// Sample test for mocked repository. WIP: not working
func TestGetOrders(t *testing.T) {
	t.Run("should return all the orders of the mocked repository", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		context := gin.New()

		repositoryMock.On("GetOrders", context).Return(ordersMock, nil)
		context.GET("/profile/orders", profile_handler.GetOrders(repositoryMock))

		req, err := http.NewRequest(http.MethodGet, "/profile/orders", bytes.NewBufferString(""))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		context.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}
