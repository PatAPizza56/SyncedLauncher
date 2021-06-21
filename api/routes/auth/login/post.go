package login

import (
	"github.com/gofiber/fiber"

	"../../../structs"
)

func Post(c *fiber.Ctx) error {
	var value string
	login := new(structs.Login)

	err := c.BodyParser(login)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Failed to parse body"))
	}

	err, stat := login.Attempt(&value)
	if err != nil {
		c.Status(stat)
		return c.Send([]byte(err.Error()))
	}

	return c.SendString(value)
}
