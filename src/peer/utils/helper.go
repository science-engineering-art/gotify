package utils

import (
	"github.com/science-engineering-art/spotify/src/peer/pb"
	"go.mongodb.org/mongo-driver/bson"
)

func ToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

func BuildQuery(req *pb.SongMetadata) *bson.M {
	query := bson.M{}

	if req.Album != nil {
		query["album"] = req.Album
	}

	if req.AlbumArtist != nil {
		query["albumartist"] = req.AlbumArtist
	}

	if req.Artist != nil {
		query["artist"] = req.Artist
	}

	if req.Comment != nil {
		query["comment"] = req.Comment
	}

	if req.Composer != nil {
		query["composer"] = req.Composer
	}

	if req.FileType != nil {
		query["filetype"] = req.FileType
	}

	if req.Format != nil {
		query["format"] = req.Format
	}

	if req.Genre != nil {
		query["genre"] = req.Genre
	}

	if req.Lyrics != nil {
		query["lyrics"] = req.Lyrics
	}

	if req.Title != nil {
		query["title"] = req.Title
	}

	if req.Year != nil {
		query["year"] = req.Year
	}

	return &query
}
