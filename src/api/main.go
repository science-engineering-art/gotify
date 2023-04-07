package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/science-engineering-art/spotify/config"
	"github.com/science-engineering-art/spotify/routes"
)

func main() {
	// Fiber instance
	app := fiber.New(fiber.Config{
		BodyLimit: 40 * 1024 * 1024, // this is the default limit of 4MB
	})

	// configure CORS
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	//run database
	config.ConnectDB()

	//routes
	routes.SongRoute(app)

	app.Listen(":5000")
}
