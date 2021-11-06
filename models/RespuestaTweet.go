package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RespuestaTweet struct {
	Id      primitive.ObjectID `bson:"_id" json":"_id,omitempty"`
	UserId  string             `bson:"userid" json":"userId,omitempty"`
	Mensaje string             `bson:"mensaje" json":"mensaje,omitempty"`
}
