package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/science-engineering-art/spotify/config"
	"github.com/science-engineering-art/spotify/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSong(c *fiber.Ctx) error {
	collection, err := config.GetMongoDbCollection(config.DbName, config.CollectionName)
	if err != nil {
		c.Status(500)
		return err
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
	}
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500)
		return err
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return err
	}

	json, _ := json.Marshal(results)
	c.Send(json)
	return nil
}

func CreateSong(c *fiber.Ctx) error {
	collection, err := config.GetMongoDbCollection(config.DbName, config.CollectionName)
	if err != nil {
		c.Status(500)
		return err
	}

	var song models.Song
	json.Unmarshal([]byte(c.Body()), &song)

	res, err := collection.InsertOne(context.Background(), song)
	if err != nil {
		c.Status(500)
		return err
	}

	response, _ := json.Marshal(res)
	c.Send(response)
	return nil
}

func UpdateSong(c *fiber.Ctx) error {
	collection, err := config.GetMongoDbCollection(config.DbName, config.CollectionName)
	if err != nil {
		c.Status(500)
		return err
	}
	var song models.Song
	json.Unmarshal([]byte(c.Body()), &song)

	update := bson.M{
		"$set": song,
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		c.Status(500)
		return err
	}

	response, _ := json.Marshal(res)
	c.Send(response)
	return nil
}

func DeleteSong(c *fiber.Ctx) error {
	collection, err := config.GetMongoDbCollection(config.DbName, config.CollectionName)

	if err != nil {
		c.Status(500)
		return err
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		c.Status(500)
		return err
	}

	jsonResponse, _ := json.Marshal(res)
	c.Send(jsonResponse)
	return nil
}
