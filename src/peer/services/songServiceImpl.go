package services

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dhowden/tag"
	"github.com/science-engineering-art/spotify/src/peer/models"
	"github.com/science-engineering-art/spotify/src/peer/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SongServiceImpl struct {
	songCollection *mongo.Collection
	ctx            context.Context
}

func NewSongService(songCollection *mongo.Collection, ctx context.Context) SongService {
	return &SongServiceImpl{songCollection, ctx}
}

func (s *SongServiceImpl) CreateSong(rawSong []byte) error {
	songBytes := bytes.NewReader(rawSong)

	m, err := tag.ReadFrom(songBytes)
	if err != nil {
		return err
	}

	objId := primitive.NewObjectID()

	// build the Song model with all its metadata
	newSong := models.Song{
		Album:       m.Album(),
		AlbumArtist: m.AlbumArtist(),
		Artist:      m.Artist(),
		Comment:     m.Comment(),
		Composer:    m.Composer(),
		FileType:    m.FileType(),
		Format:      m.Format(),
		Genre:       m.Genre(),
		Id:          objId,
		RawSong:     rawSong,
		Lyrics:      m.Lyrics(),
		Title:       m.Title(),
		Year:        m.Year(),
	}

	res, err := s.songCollection.InsertOne(s.ctx, newSong)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return errors.New("post with that title already exists")
		}
		return err
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"title": 1}, Options: opt}

	if _, err := s.songCollection.Indexes().CreateOne(s.ctx, index); err != nil {
		return errors.New("could not create index for title")
	}

	query := bson.M{"_id": res.InsertedID}
	if err = s.songCollection.FindOne(s.ctx, query).Decode(&newSong); err != nil {
		return err
	}

	return nil
}

func (s *SongServiceImpl) UpdateSong(data *models.UpdateSongRequest) error {
	doc, err := utils.ToDoc(data)
	if err != nil {
		return err
	}

	query := bson.D{{Key: "_id", Value: data.Id}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := s.songCollection.FindOneAndUpdate(s.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedPost *models.Song
	if err := res.Decode(&updatedPost); err != nil {
		return errors.New("no post with that Id exists")
	}

	return nil
}

func (s *SongServiceImpl) GetSongById(id string) (*models.Song, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var song *models.Song

	if err := s.songCollection.FindOne(s.ctx, query).Decode(&song); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return song, nil
}

func (s *SongServiceImpl) GetSongs() ([]*models.Song, error) {

	opt := options.FindOptions{}
	// opt.SetLimit(int64(limit))
	// opt.SetSkip(int64(skip))
	// opt.SetSort(bson.M{"created_at": -1})
	query := bson.M{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.songCollection.Find(ctx, query, &opt)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer cursor.Close(s.ctx)

	var songs []*models.Song

	for cursor.Next(s.ctx) {
		song := &models.Song{}
		err := cursor.Decode(song)

		if err != nil {
			return nil, err
		}

		songs = append(songs, song)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(songs) == 0 {
		return []*models.Song{}, nil
	}

	return songs, nil
}

func (s *SongServiceImpl) DeleteSong(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := s.songCollection.DeleteOne(s.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}
