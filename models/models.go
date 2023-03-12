package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateMusicRequest struct {
	Title     string    `json:"title,omitempty" bson:"title,omitempty"`
	Album     string    `json:"album,omitempty" bson:"album,omitempty"`
	Gender    string    `json:"gender,omitempty" bson:"gender,omitempty"`
	Author    string    `json:"author,omitempty" bson:"author,omitempty"`
	Location  string    `json:"location,omitempty" bson:"location,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type DBMusic struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Album     string             `json:"album,omitempty" bson:"album,omitempty"`
	Gender    string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Author    string             `json:"author,omitempty" bson:"author,omitempty"`
	Location  string             `json:"location,omitempty" bson:"location,omitempty"`
	CreateAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UpdateMusic struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Album     string             `json:"album,omitempty" bson:"album,omitempty"`
	Gender    string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Author    string             `json:"author,omitempty" bson:"author,omitempty"`
	Location  string             `json:"location,omitempty" bson:"location,omitempty"`
	CreateAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
