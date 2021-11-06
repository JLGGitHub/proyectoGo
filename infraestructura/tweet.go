package infraestructura

import (
	"context"
	"log"
	"time"

	"ejercicios/cursoUdemy/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertarTweet(t models.GraboTweet) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConexion.Database("twiter")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserId,
		"mensaje": t.Mensaje,
	}
	resultado, err := col.InsertOne(ctx, registro)

	if err != nil {

		return "", false, err
	}

	objId, _ := resultado.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil
}

func LeerTweet(id string, pagina int64) ([]*models.RespuestaTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConexion.Database("twiter")
	col := db.Collection("tweet")

	var resultados []*models.RespuestaTweet

	condicion := bson.M{
		"userid": id,
	}
	opc := options.Find()
	opc.SetLimit(10)
	opc.SetSkip((pagina - 1) * 10)

	cursor, err := col.Find(ctx, condicion, opc)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.RespuestaTweet
		err := cursor.Decode(&registro)
		if err != nil {

			return resultados, false
		}
		resultados = append(resultados, &registro)
	}

	return resultados, true
}
