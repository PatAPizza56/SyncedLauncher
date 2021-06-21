package games

import (
	"strconv"

	"github.com/gofiber/fiber"

	"../../structs"
)

func Post(c *fiber.Ctx) error {
	var id int
	game := new(structs.Game)
	var user structs.User

	err := c.BodyParser(game)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Failed to parse body"))
	}

	err, stat := user.GetByToken(c.Query("token"))
	if err != nil {
		c.Status(stat)
		return err
	}

	game.UserID = user.ID

	err, stat = game.Post(&id)
	if err != nil {
		c.Status(stat)
		return c.Send([]byte(err.Error()))
	}

	return c.Send([]byte(strconv.Itoa(id)))
}
