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

	jsonFile, err := os.Open("./initial_data/users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// Leer el archivo de clientes y meterlo en clients[]
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var clients []models.Client
	json.Unmarshal(byteValue, &clients)

	// Inserción de los clientes en la colección "user"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < len(clients); i++ {
		insertResult, err := db.Collections["user"].InsertOne(ctx, clients[i])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(insertResult.InsertedID)
	}
}
