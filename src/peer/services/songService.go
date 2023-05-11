package services

import "github.com/science-engineering-art/spotify/src/peer/models"

type SongService interface {
	CreateSong([]byte) error
	UpdateSong(*models.UpdateSongRequest) error
	GetSongById(string) (*models.Song, error)
	GetSongs() ([]*models.Song, error)
	DeleteSong(string) error
}
