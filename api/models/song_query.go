package models

type SongQuery struct {
	Artist string `json:"artist,omitempty" bson:"artist,omitempty"`
	Album  string `json:"album,omitempty" bson:"album,omitempty"`
	Genre  string `json:"genre,omitempty" bson:"genre,omitempty"`
	Title  string `json:"title,omitempty" bson:"title,omitempty"`
}
