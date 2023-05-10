package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SongDTO struct {
	Artist string             `json:"artist,omitempty"`
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `json:"title,omitempty"`
	Year   int                `json:"year,omitempty"`
}
