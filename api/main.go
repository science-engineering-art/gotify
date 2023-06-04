package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
	"github.com/science-engineering-art/gotify/api/net"
	"github.com/science-engineering-art/gotify/api/routes"
)

func main() {
	// Fiber instance
	app := fiber.New(fiber.Config{
		// the limit is 1Gb
		BodyLimit: 1024 * 1024 * 1024,
	})

	// Configuring the public folder to serve static files
	app.Static("/", "./public", fiber.Static{
		// Enable gzip compression and response to HTTP Range Requests
		Compress: true,
	})

	// Configure CORS
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Use(func(c *fiber.Ctx) error {
		// Generate a unique identifier
		id := uuid.New()

		// Set the unique identifier in the context of the request
		c.Context().SetUserValue("requestId", id)

		// Continue with the application
		return c.Next()
	})

	// Configure the routes
	routes.SongRoute(app)

	go net.Broadcast(53123)

	// Enable port for listening
	app.Listen("0.0.0.0:80")
}
