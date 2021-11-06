package models

type GraboTweet struct {
	UserId  string `bson:"userid" json:"userid,omitempty"`
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
}
