package models

import (
	"github.com/dhowden/tag"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Song struct {
	Album       string             `json:"album,omitempty" bson:"album,omitempty"`
	AlbumArtist string             `json:"albumartist,omitempty" bson:"albumartist,omitempty"`
	Artist      string             `json:"artist,omitempty" bson:"artist,omitempty"`
	Comment     string             `json:"comment,omitempty" bson:"comment,omitempty"`
	Composer    string             `json:"composer,omitempty" bson:"composer,omitempty"`
	FileType    tag.FileType       `json:"filetype,omitempty" bson:"filetype,omitempty"`
	Format      tag.Format         `json:"format,omitempty" bson:"format,omitempty"`
	Genre       string             `json:"genre,omitempty" bson:"genre,omitempty"`
	Id          primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Lyrics      string             `json:"lyrics,omitempty" bson:"lyrics,omitempty"`
	RawSong     string             `json:"rawsong" bson:"rawsong" binding:"required"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Year        int                `json:"year,omitempty" bson:"year,omitempty"`
}
