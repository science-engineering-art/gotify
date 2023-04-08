package controllers

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dhowden/tag"
	"github.com/science-engineering-art/spotify/config"
	"github.com/science-engineering-art/spotify/models"
	"github.com/science-engineering-art/spotify/responses"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var songCollection *mongo.Collection = config.GetCollection(config.DB, "songs")
var validate = validator.New()

func CreateSong(c *fiber.Ctx) error {

	fileForm, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
		return err
	}

	filename := fileForm.Filename

	err = c.SaveFile(fileForm, fmt.Sprintf("./%s", filename))
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()

	buffer := make([]byte, size)

	file.Read(buffer)

	songBytes := bytes.NewReader(buffer)
	m, err := tag.ReadFrom(songBytes)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId := primitive.NewObjectID()

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
		Lyrics:      m.Lyrics(),
		RawSong:     base64.RawStdEncoding.EncodeToString(buffer),
		Title:       m.Title(),
		Year:        m.Year(),
	}

	objID, err := songCollection.InsertOne(ctx, newSong)
	if err != nil {
		return c.Status(
			http.StatusInternalServerError).JSON(
			responses.SongResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			},
		)
	}

	return c.Status(
		http.StatusCreated).JSON(
		responses.SongResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    &fiber.Map{"_id": objID},
		},
	)
}

func GetSong(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	songId := c.Params("songId")
	var song models.Song
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(songId)

	err := songCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&song)
	if err != nil {
		return c.Status(
			http.StatusInternalServerError).JSON(
			responses.SongResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.SongResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    &fiber.Map{"data": song},
		},
	)
}

func EditSong(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	songId := c.Params("songId")
	var song models.Song
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(songId)

	//validate the request body
	if err := c.BodyParser(&song); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.SongResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			},
		)
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&song); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.SongResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"tile": song.Title, "album": song.Album, "Genre": song.Genre}

	result, err := songCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.SongResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	//get updated user details
	var updatedSong models.Song
	if result.MatchedCount == 1 {
		err := songCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedSong)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.SongResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.SongResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedSong}})
}

func DeleteSong(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	songId := c.Params("songId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(songId)

	result, err := songCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.SongResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.SongResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.SongResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
	)
}

func GetSongs(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var songs []models.Song
	defer cancel()

	results, err := songCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.SongResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleSong models.Song
		if err = results.Decode(&singleSong); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.SongResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		songs = append(songs, singleSong)
	}

	return c.Status(http.StatusOK).JSON(
		responses.SongResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": songs}},
	)
}
