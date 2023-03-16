package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/science-engineering-art/spotify/config"
	"github.com/science-engineering-art/spotify/controllers"
)

func main() {
	app := fiber.New()

	app.Get("/songs/:id?", controllers.GetSong)
	app.Post("/songs", controllers.CreateSong)
	app.Put("/songs/:id", controllers.UpdateSong)
	app.Delete("/songs/:id", controllers.DeleteSong)

	app.Listen(config.Port)
}
