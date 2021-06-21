package users

import (
	"strings"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"

	"../../structs"
	"../../utils"
)

func Put(c *fiber.Ctx) error {
	user := new(structs.User)
	var aUser structs.User

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

	err = c.BodyParser(user)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Failed to parse body"))
	}

	if len(user.Username) < 2 {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Username must be at least 2 characters long"))
	} else if !strings.Contains(user.Email, "@") {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Please enter a valid email"))
	} else if len(user.Password) < 8 {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Password must be at least 8 characters long"))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.Status(fiber.StatusConflict)
		return c.Send([]byte("Failed to hash password"))
	}
	user.Password = string(hashedPassword)

	err, stat = user.Put(utils.Method(c.Query("method")), utils.Value(c.Params("value")))
	if err != nil {
		c.Status(stat)
		return c.Send([]byte(err.Error()))
	}

	return c.Send([]byte("Success"))
}
