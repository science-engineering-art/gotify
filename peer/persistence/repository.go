package persistence

import (
	"github.com/science-engineering-art/spotify/src/peer/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SongRepository interface {
	CreateSong(key string, rawSong *[]byte) error
	GetSongById(objID *primitive.ObjectID) (*models.Song, error)
	UpdateSong(objID *primitive.ObjectID, updatedSong *bson.M) error
	RemoveSongById(objID *primitive.ObjectID) error
	FilterSongs(query *bson.M) ([]*models.Song, error)
}
