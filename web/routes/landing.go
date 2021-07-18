package routes

import (
	"bytes"
	"html/template"

	"github.com/gofiber/fiber"
)

func Landing(c *fiber.Ctx) error {
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
}
