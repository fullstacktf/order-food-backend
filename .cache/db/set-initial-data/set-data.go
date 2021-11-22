package main // si le pongo otro nombre de paquete peta al hacer go run. No lo entiendo :(

import (
	"comiditapp/api/src/database"
	"comiditapp/api/src/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	db := database.GetDB()
	db.DropDB()

	// Lectura de usuarios cliente
	var clients []models.User
	byteValue := readData("./initial_data/clients.json")
	json.Unmarshal(byteValue, &clients)

	// Lectura de usuarios restaurante
	var restaurants []models.User
	byteValue = readData("./initial_data/restaurants.json")
	json.Unmarshal(byteValue, &restaurants)

	// Lectura de orders
	var orders []models.Order
	byteValue = readData("./initial_data/orders.json")
	json.Unmarshal(byteValue, &orders)

	// INSERCIONES EN LA BBDD

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Inserción de clients
	for i := 0; i < len(clients); i++ {
		_, err := db.Collections["user"].InsertOne(ctx, clients[i])
		if err != nil {
			fmt.Println(err)
		}
	}

	// Inserción de restaurantes
	for i := 0; i < len(restaurants); i++ {
		_, err := db.Collections["user"].InsertOne(ctx, restaurants[i])
		if err != nil {
			fmt.Println(err)
		}
	}

	// Inserción de orders
	for i := 0; i < len(orders); i++ {
		_, err := db.Collections["order"].InsertOne(ctx, orders[i])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func readData(fileName string) []byte {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}
