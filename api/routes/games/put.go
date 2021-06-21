package games

import (
	"github.com/gofiber/fiber"

	"../../structs"
	"../../utils"
)

func Put(c *fiber.Ctx) error {
	game := new(structs.Game)
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

	err = c.BodyParser(game)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Failed to parse body"))
	}

	game.UserID = user.ID

	err, stat = game.Put(utils.Method(c.Query("method")), utils.Value(c.Params("value")))
	if err != nil {
		c.Status(stat)
		return c.Send([]byte(err.Error()))
	}

	return c.Send([]byte("Success"))
}
