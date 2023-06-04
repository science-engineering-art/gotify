package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/science-engineering-art/gotify/api/controllers"
)

func SongRoute(app *fiber.App) {
	//All routes related to songs comes here
	app.Post("/api/song", controllers.CreateSong)
	// app.Get("/api/song/:songId", controllers.GetSongById)
	// app.Put("/api/song/:songId", controllers.UpdateSong)
	// app.Delete("/api/song/:songId", controllers.RemoveSongById)
	// app.Post("/api/songs", controllers.FilterSongs)
}
