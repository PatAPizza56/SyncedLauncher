package server

import (
	"bytes"
	"html/template"

	"github.com/gofiber/fiber"

	routes "../routes"
)

func Start() {
	app := fiber.New()

	app.Static("/static", "./src/static")

	app.Get("/", routes.Home)
	app.Get("/games/:title", routes.Games)

	app.Use(func(c *fiber.Ctx) error {
		page, err := template.ParseFiles("src/html/index.html")
		if err != nil {
			return err
		}

		var template bytes.Buffer
		err = page.Execute(&template, nil)
		if err != nil {
			return err
		}

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(template.String())
	})

	app.Listen(":3000")
}
