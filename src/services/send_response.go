package services

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int
	Message []string
	Error   []string
	Data    map[string]interface{}
}

func SendResponse(c *gin.Context, response Response) {
	if len(response.Message) > 0 {
		c.JSON(response.Status, map[string]interface{}{"message": strings.Join(response.Message, "; ")})
	} else if len(response.Data) > 0 {
		c.JSON(response.Status, response.Data)
	} else if len(response.Error) > 0 {
		c.JSON(response.Status, map[string]interface{}{"error": strings.Join(response.Error, "; ")})
	}
}
