package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Song struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Title    string             `json:"title,omitempty" validate:"required"`
	Genre    string             `json:"genre,omitempty" validate:"required"`
	Author   string             `json:"author,omitempty" validate:"required"`
	Album    string             `json:"album,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
}
