package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre    string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos string             `bson:"apellidos" json:"apellidos,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Pass      string             `bson:"pass" json:"pass,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
}
