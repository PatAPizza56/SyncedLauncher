package games

import (
	"github.com/gofiber/fiber"

	"../../structs"
	"../../utils"
)

func Get(c *fiber.Ctx) error {
	var game structs.Game

	err, stat := game.Get(utils.Method(c.Query("method")), utils.Value(c.Params("value")))
	if err != nil {
		c.Status(stat)
		return err
	}

	return c.JSON(game)
}
