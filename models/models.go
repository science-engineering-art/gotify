package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Song struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitempty"`
	Album    string             `json:"album,omitempty" bson:"album,omitempty"`
	Gender   string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Author   string             `json:"author,omitempty" bson:"author,omitempty"`
	Location string             `json:"location,omitempty" bson:"location,omitempty"`
}
