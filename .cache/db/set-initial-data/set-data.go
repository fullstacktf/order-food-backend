package main // si le pongo otro nombre de paquete peta al hacer go run. No lo entiendo :(

import (
	"comiditapp/api/src/database"
	"comiditapp/api/src/env"
	"comiditapp/api/src/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const dataDirectory = "./.cache/db/set-initial-data/initial-data"

func main() {

	env.LoadEnv()

	db := database.GetDB()
	db.DropDB()

	// Lectura de usuarios cliente
	var clients []models.User
	byteValue := readData(dataDirectory + "/clients.json")
	json.Unmarshal(byteValue, &clients)

	var interfaceClients []interface{} = make([]interface{}, len(clients))
	for index, client := range clients {
		interfaceClients[index] = client
	}

	// Lectura de usuarios restaurante
	var restaurants []models.User
	byteValue = readData(dataDirectory + "/restaurants.json")
	json.Unmarshal(byteValue, &restaurants)

	var interfaceRestaurants []interface{} = make([]interface{}, len(restaurants))
	for index, restaurant := range restaurants {
		interfaceRestaurants[index] = restaurant
	}

	// Lectura de orders
	var orders []models.Order
	byteValue = readData(dataDirectory + "/orders.json")
	json.Unmarshal(byteValue, &orders)

	var interfaceOrders []interface{} = make([]interface{}, len(orders))
	for index, order := range orders {
		interfaceOrders[index] = order
	}

	// INSERCIONES EN LA BBDD

	insertMany(interfaceClients, db.Collections["user"])
	insertMany(interfaceRestaurants, db.Collections["user"])
	insertMany(interfaceOrders, db.Collections["order"])
}

func insertMany(interfaceSlice []interface{}, collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertMany(ctx, interfaceSlice)
	if err != nil {
		log.Fatal(err)
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
