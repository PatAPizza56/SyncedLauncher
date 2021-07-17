package users

import (
	"strconv"

	"github.com/gofiber/fiber"

	"../../structs"
	"../../utils"
)

func Delete(c *fiber.Ctx) error {
	var game structs.Game
	var user structs.User
	var aUser structs.User
	var token structs.Token

	err, stat := user.Get(utils.Method(c.Query("method")), utils.Value(c.Params("value")))
	if err != nil {
		c.Status(stat)
		return err
	}

	err, stat = aUser.GetByToken(c.Query("token"))
	if err != nil {
		c.Status(stat)
		return err
	}

	if user.ID != aUser.ID {
		c.Status(fiber.StatusUnauthorized)
		return c.Send([]byte("Token not valid"))
	}

	err, stat = game.Delete("UserID", strconv.Itoa(user.ID))
	if err != nil {
		c.Status(stat)
		return err
	}

	err, stat = user.Delete(utils.Method(c.Query("method")), utils.Value(c.Params("value")))
	if err != nil {
		c.Status(stat)
		return err
	}

	err, stat = token.Delete("UserID", strconv.Itoa(user.ID))
	if err != nil {
		c.Status(stat)
		return err
	}

	return c.Send([]byte("Success"))
}
