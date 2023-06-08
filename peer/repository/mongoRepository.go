package repository

import (
	"github.com/science-engineering-art/gotify/peer/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoRepository interface {
	CreateSong(key string, rawSong *[]byte) error
	GetSongById(objID *primitive.ObjectID) (*models.Song, error)
	UpdateSong(objID *primitive.ObjectID, updatedSong *bson.M) error
	RemoveSongById(objID *primitive.ObjectID) error
	SongFilter(query *bson.M) ([]*models.Song, error)
}
