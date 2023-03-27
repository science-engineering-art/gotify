package controllers

import (
	"context"
	"net/http"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var song models.Song
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&song); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.SongResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&song); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.SongResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newSong := models.Song{
		Id:       primitive.NewObjectID(),
		Title:    song.Title,
		Location: song.Location,
		Genre:    song.Genre,
		Album:    song.Album,
		Author:   song.Author,
	}

	result, err := songCollection.InsertOne(ctx, newSong)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.SongResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.SongResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetSong(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	songId := c.Params("songId")
	var song models.Song
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(songId)

	err := songCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&song)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.SongResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.SongResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": song}})
}

func EditSong(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	songId := c.Params("songId")
	var song models.Song
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(songId)

	//validate the request body
	if err := c.BodyParser(&song); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.SongResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&song); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.SongResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"tile": song.Title, "location": song.Location, "author": song.Author, "album": song.Album, "Genre": song.Genre}

	result, err := songCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

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
