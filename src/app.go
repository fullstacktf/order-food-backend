package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/", getHandler)

	if err := r.Run(":3000"); err != nil {
		log.Fatal(err.Error())
	}
}

func getHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "You are at comiditapp homepage🥳🥳",
	})
}
