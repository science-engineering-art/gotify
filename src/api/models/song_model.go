package models

import (
	"github.com/dhowden/tag"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Song struct {
	Album       string             `json:"album,omitempty"`
	AlbumArtist string             `json:"albumartist,omitempty"`
	Artist      string             `json:"artist,omitempty"`
	Comment     string             `json:"comment,omitempty"`
	Composer    string             `json:"composer,omitempty"`
	FileType    tag.FileType       `json:"filetype,omitempty"`
	Format      tag.Format         `json:"format,omitempty"`
	Genre       string             `json:"genre,omitempty"`
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Lyrics      string             `json:"lyrics,omitempty"`
	RawSong     string             `json:"rawsong,omitempty"`
	Title       string             `json:"title,omitempty"`
	Year        int                `json:"year,omitempty"`
}
