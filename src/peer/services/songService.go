package services

import (
	"github.com/science-engineering-art/spotify/src/peer/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SongService interface {
	CreateSong([]byte) error
	UpdateSong(*models.UpdatedSong) error
	GetSongById(primitive.ObjectID) (*models.Song, error)
	FilterSongs(*bson.M) ([]*models.Song, error)
	RemoveSongById(primitive.ObjectID) error
}
