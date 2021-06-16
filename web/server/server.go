package server

import (
	"github.com/gofiber/fiber"

	routes "../routes"
)

func Start() {
	app := fiber.New()

	app.Static("/static", "./src/static")

	app.Get("/", routes.Home)
	app.Get("/games/:title", routes.Games)

	app.Listen(":3000")
}
