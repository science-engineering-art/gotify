package persistence

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDb struct {
	repository SongRepository
}

func NewMongoDb(database, collection, mongoDbIP string) (s *MongoDb) {
	mongoDbUri := fmt.Sprintf("mongodb://user:password@%s:27017/?maxPoolSize=20&w=majority", mongoDbIP)
	fmt.Println("Trying to connect...", mongoDbIP)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDbUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDb successfully connected to", mongoDbIP)

	// Collections
	songCollection := client.Database(database).Collection(collection)

	songRepo := NewSongRepository(songCollection)

	newMongo := MongoDb{}
	newMongo.repository = songRepo

	return &newMongo
}

func (s *MongoDb) Create(key []byte, data *[]byte) error {
	b64 := base64.RawStdEncoding.EncodeToString(key)

	err := s.repository.CreateSong(b64, data)
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoDb) Read(key []byte) (data *[]byte, err error) {
	b64 := base64.RawStdEncoding.EncodeToString(key)

	objID, err := primitive.ObjectIDFromHex(b64)
	if err != nil {
		return nil, err
	}

	song, err := s.repository.GetSongById(&objID)
	if err != nil {
		return nil, err
	}

	return &song.RawSong, nil
}

func (s *MongoDb) Delete(key []byte) error {
	b64 := base64.RawStdEncoding.EncodeToString(key)

	objID, err := primitive.ObjectIDFromHex(b64)
	if err != nil {
		return err
	}

	s.repository.RemoveSongById(&objID)
	if err != nil {
		return err
	}

	return nil
}
