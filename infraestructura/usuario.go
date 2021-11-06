package infraestructura

import (
	"context"
	"ejercicios/cursoUdemy/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscarUsuario(id string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConexion.Database("twiter")
	col := db.Collection("usuarios")

	var perfil models.Usuario

	objId, _ := primitive.ObjectIDFromHex(id)
	condicion := bson.M{
		"_id": objId,
	}
	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Pass = ""
	if err != nil {
		fmt.Println("No se encontro el registro" + err.Error())
		return perfil, err
	}
	return perfil, nil
}

func Insertar(item models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	db := MongoConexion.Database("twiter")
	col := db.Collection("usuarios")

	item.Pass, _ = EncriptarPass(item.Pass)

	_, err := col.InsertOne(ctx, item)

	if err != nil {
		return "", false, err
	}
	return "", true, nil
}

func Modificar(item models.Usuario, Id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	fmt.Printf("****** %v ", Id)
	defer cancel()
	db := MongoConexion.Database("twiter")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{}, 0)

	if len(item.Nombre) > 0 {
		registro["nombre"] = item.Nombre
	}
	if len(item.Apellidos) > 0 {
		registro["apellidos"] = item.Apellidos
	}
	if len(item.Email) > 0 {
		registro["email"] = item.Email
	}
	if len(item.Pass) > 0 {
		registro["pass"] = item.Pass
	}

	updateString := bson.M{
		"$set": registro,
	}

	objId, _ := primitive.ObjectIDFromHex(Id)

	filtro := bson.M{"_id": bson.M{"$eq": objId}}

	_, err := col.UpdateOne(ctx, filtro, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
