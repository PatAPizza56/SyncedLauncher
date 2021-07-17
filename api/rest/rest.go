package rest

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"

	"../routes/auth/login"
	"../routes/auth/register"
	"../routes/games"
	"../routes/stripe/connect"
	"../routes/user"
	"../routes/users"
)

func Start() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET,POST,HEAD, OPTIONS, PUT, DELETE, PATCH",
		AllowHeaders:  "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		ExposeHeaders: "Origin",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Synced Studios API")
	})

	app.Post("/games", games.Post)
	app.Get("/games/:value", games.Get)
	app.Put("/games/:value", games.Put)
	app.Delete("/games/:value", games.Delete)

	app.Get("/user/:value", user.Get)

	app.Get("/users/:value", users.Get)
	app.Put("/users/:value", users.Put)
	app.Delete("/users/:value", users.Delete)

	app.Get("/stripe/connect/:value", connect.Get)

	app.Post("/auth/login", login.Post)
	app.Post("/auth/register", register.Post)

	app.Listen(":8000")
}
