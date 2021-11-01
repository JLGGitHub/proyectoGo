package infraestructura

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConexion = ConectarDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.xogpy.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	fmt.Println("Conexion Exitosa!!! ")
	return client
}

func ChequeoConexion() int {
	err := MongoConexion.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
