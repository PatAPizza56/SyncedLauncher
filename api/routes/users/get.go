package users

import (
	"github.com/gofiber/fiber"

	"../../structs"
	"../../utils"
)

func Get(c *fiber.Ctx) error {
	var user structs.User

	err, stat := user.Get(utils.Method(c.Query("method")), utils.Value(c.Params("value")))
	if err != nil {
		c.Status(stat)
		return err
	}

	user.Password = ""

	return c.JSON(user)
}
