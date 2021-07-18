package server

import (
	"bytes"
	"html/template"

	"github.com/gofiber/fiber"

	"../routes"
)

func Start() {
	app := fiber.New()

	app.Static("/static", "./src/static")

	//app.Get("/", routes.Home)
	//app.Get("/games/:title", routes.Games)
	app.Get("/", routes.Landing)
	app.Get("/register", routes.Register)
	app.Get("/register/success", routes.RegisterSuccess)

	app.Use(func(c *fiber.Ctx) error {
		page, err := template.ParseFiles("src/html/404/index.html")
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
