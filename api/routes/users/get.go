package users

import (
	"github.com/gofiber/fiber"

	"../../structs"
	"../../utils"
)

type UserResponse struct {
	ID       int
	Username string
	PfpURL   string
}

func Get(c *fiber.Ctx) error {
	var user structs.User

	err, stat := user.Get(utils.Method(c.Query("method")), utils.Value(c.Params("value")))
	if err != nil {
		c.Status(stat)
		return err
	}

	return c.JSON(&UserResponse{ID: user.ID, Username: user.Username, PfpURL: user.PfpURL})
}
