package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/science-engineering-art/spotify/controllers"
)

func SongRoute(app *fiber.App) {
	//All routes related to songs comes here
	app.Post("/song", controllers.CreateSong)
	app.Get("/song/:songId", controllers.GetSong)
	app.Put("/song/:songId", controllers.EditSong)
	app.Delete("/song/:songId", controllers.DeleteSong)
	app.Get("/songs", controllers.GetSongs)
}
