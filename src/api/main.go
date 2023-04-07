package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/science-engineering-art/spotify/config"
	"github.com/science-engineering-art/spotify/routes"
)

func main() {
	app := fiber.New()

	//run database
	config.ConnectDB()

	//routes
	routes.SongRoute(app)

	app.Listen(":5000")
}
