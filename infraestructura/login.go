package infraestructura

import (
	"context"
	"ejercicios/cursoUdemy/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, pass string) (models.Usuario, bool) {

	usuario, exite, _ := ChequeoYaExisteUsuario(email)
	fmt.Println(usuario.Nombre)
	if !exite {
		return usuario, false
	}

	pass2 := []byte(pass)
	passinfraestructura := []byte(usuario.Pass)

	err := bcrypt.CompareHashAndPassword(passinfraestructura, pass2)
	if err != nil {
		return usuario, false
	}

	return usuario, true
}

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	fmt.Printf("Chequeo de usuario->  %s\n", email)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConexion.Database("twiter")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}
	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	Id := resultado.Id.Hex()
	if err != nil {
		fmt.Println(err.Error())
		return resultado, false, Id
	}
	return resultado, true, Id
}

func EncriptarPass(pass string) (string, error) {

	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
