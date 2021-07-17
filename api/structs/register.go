package structs

import (
	"errors"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"

	"../utils"
)

type Register struct {
	FullName string
	Username string
	Email    string
	Password string
}

func (register *Register) Attempt(newUserID *int, newUserToken *string) (error, int) {
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

	user = User{FullName: register.FullName, Username: register.Username, Email: register.Email, Password: string(hashedPassword), PfpURL: "https://cdn.discordapp.com/attachments/814608238917451776/860550308986880020/unknown.png"}

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

	*newUserID = userID
	*newUserToken = token.Value

	return nil, 0
}
