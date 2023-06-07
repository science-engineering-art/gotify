package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/science-engineering-art/gotify/api/controllers"
)

func SongRoute(app *fiber.App) {
	//All routes related to songs comes here
	app.Post("/song", controllers.CreateSong)
	// app.Get("/song/:songId", controllers.GetSongById)
	// app.Put("/song/:songId", controllers.UpdateSong)
	// app.Delete("/song/:songId", controllers.RemoveSongById)
	app.Post("/songs", controllers.FilterSongs)
}
