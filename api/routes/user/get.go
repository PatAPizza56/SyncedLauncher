package user

import (
	"strconv"

	"github.com/gofiber/fiber"

	"../../structs"
)

func Get(c *fiber.Ctx) error {
	var token structs.Token
	var user structs.User

	err, stat := token.Get("Value", c.Params("value"))
	if err != nil {
		c.Status(stat)
		return err
	}

	err, stat = user.Get("ID", strconv.Itoa(token.UserID))
	if err != nil {
		c.Status(stat)
		return err
	}

	user.Password = ""

	return c.JSON(user)
}
