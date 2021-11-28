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
)

var repositoryMock repository.MockedOrdersRepository

var ordersMock []models.Order = []models.Order{
	{Id: "1", RestaurantId: "1", ClientId: "1", Status: enums.Ordered, TotalPrice: 30.30, Products: []models.ProductInfo{{ProductId: "1", Quantity: 1}}},
	{Id: "2", RestaurantId: "1", ClientId: "2", Status: enums.Preparing, TotalPrice: 45.30, Products: []models.ProductInfo{{ProductId: "3", Quantity: 2}, {ProductId: "1", Quantity: 2}}},
	{Id: "3", RestaurantId: "2", ClientId: "2", Status: enums.Preparing, TotalPrice: 45.30, Products: []models.ProductInfo{{ProductId: "3", Quantity: 2}, {ProductId: "1", Quantity: 2}}},
	{Id: "4", RestaurantId: "1", ClientId: "2", Status: enums.Preparing, TotalPrice: 45.30, Products: []models.ProductInfo{{ProductId: "3", Quantity: 2}, {ProductId: "1", Quantity: 2}}},
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
