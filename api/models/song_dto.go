package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SongDTO struct {
	Artist string             `json:"artist,omitempty" bson:"artist,omitempty"`
	Id     primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Title  string             `json:"title,omitempty" bson:"title,omitempty"`
	Year   int                `json:"year,omitempty" bson:"year,omitempty"`
}
