package structs

import (
	"errors"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"

	"../utils"
)

type Register struct {
	Username string
	Email    string
	Password string
}

func (register *Register) Attempt(value *string) (error, int) {
	var user User
	var userID int

	var token Token
	var tokenID int

	err1, _ := user.Get("Username", register.Username)
	err2, _ := user.Get("Email", register.Email)

	if err1 == nil && err2 == nil {
		return errors.New("Username or Email is already taken"), fiber.StatusBadRequest
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), 10)
	if err != nil {
		return errors.New("Failed to hash password"), fiber.StatusConflict
	}

	user = User{Username: register.Username, Email: register.Email, Password: string(hashedPassword), PFPURL: "https://cdn.discordapp.com/avatars/538009668035805195/a491901ebe4517dced854b7beb46341d.webp?size=128"}

	err, stat := user.Post(&userID)
	if err != nil {
		return err, stat
	}

	generatedToken, err := utils.GenerateToken(16)
	if err != nil {
		return errors.New("Failed to generate token"), fiber.StatusConflict
	}

	token = Token{UserID: userID, Value: generatedToken}

	err, stat = token.Post(&tokenID)
	if err != nil {
		return err, stat
	}

	*value = token.Value

	return nil, 0
}
