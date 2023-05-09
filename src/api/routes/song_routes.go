package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/science-engineering-art/spotify/controllers"
)

func SongRoute(app *fiber.App) {
	//All routes related to songs comes here
	app.Post("/api/song", controllers.CreateSong)
	app.Get("/api/song/:songId", controllers.GetSong)
	app.Put("/api/song/:songId", controllers.EditSong)
	app.Delete("/api/song/:songId", controllers.DeleteSong)
	app.Get("/api/songs", controllers.GetSongs)
	app.Get("/api/greet-peer", controllers.GreetPeer)
}
