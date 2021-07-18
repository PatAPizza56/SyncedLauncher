package register

import (
	"strings"

	"github.com/gofiber/fiber"

	"../../../structs"
)

type RegisterResponse struct {
	UserID    int
	UserToken string
}

func Post(c *fiber.Ctx) error {
	var newUserID int
	var newUserToken string

	register := new(structs.Register)

	err := c.BodyParser(register)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Failed to parse body"))
	}

	if len(register.Username) < 2 {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Username must be at least 2 characters long"))
	} else if len(register.Email) < 2 {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Email must be at least 2 characters long"))
	} else if !strings.Contains(register.Email, "@") {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Please enter a valid email"))
	} else if len(register.Password) < 8 {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Password must be at least 8 characters long"))
	}

	err, stat := register.Attempt(&newUserID, &newUserToken)
	if err != nil {
		c.Status(stat)
		return c.Send([]byte(err.Error()))
	}

	return c.JSON(&RegisterResponse{UserID: newUserID, UserToken: newUserToken})
}
