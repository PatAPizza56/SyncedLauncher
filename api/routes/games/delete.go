package games

import (
	"github.com/gofiber/fiber"

	"../../structs"
	"../../utils"
)

func Delete(c *fiber.Ctx) error {
	var game structs.Game
	var user structs.User

	err, stat := game.Get(utils.Method(c.Query("method")), utils.Value(c.Params("value")))
	if err != nil {
		c.Status(stat)
		return err
	}

	err, stat = user.GetByToken(c.Query("token"))
	if err != nil {
		c.Status(stat)
		return err
	}

	if game.UserID != user.ID {
		c.Status(fiber.StatusUnauthorized)
		return c.Send([]byte("Token not valid"))
	}

	err, stat = game.Delete(utils.Method(c.Query("method")), utils.Value(c.Params("value")))
	if err != nil {
		c.Status(stat)
		return err
	}

	return c.Send([]byte("Success"))
}
