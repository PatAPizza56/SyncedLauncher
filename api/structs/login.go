package structs

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber"

	"../utils"
)

type Login struct {
	Username string
	Password string
}

func (login *Login) Attempt(value *string) (error, int) {
	var user User
	var token Token

	err1, _ := user.Get("Username", login.Username)
	if err1 != nil || utils.CompareHash(user.Password, login.Password) != nil {
		return errors.New("Failed to login, recheck all credentials"), fiber.StatusBadRequest
	}

	generatedToken, err := utils.GenerateToken(16)
	if err != nil {
		return errors.New("Failed to generate token"), fiber.StatusConflict
	}

	token = Token{UserID: user.ID, Value: generatedToken}

	err, stat := token.Put("UserID", strconv.Itoa(user.ID))
	if err != nil {
		return err, stat
	}

	*value = token.Value

	return nil, 0
}
