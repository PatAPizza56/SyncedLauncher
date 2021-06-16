package games

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber"

	database "../../database"
	structs "../../structs"
)

func Put(c *fiber.Ctx) error {
	user := new(structs.User)

	err := c.BodyParser(user)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Failed to parse body"))
	}

	status, message := update(c.Query("method"), strings.ReplaceAll(c.Params("value"), "%20", " "), *user)

	c.Status(status)
	return c.Send([]byte(message))
}

func update(method string, value string, user structs.User) (int, string) {
	var request string

	username := user.Username
	email := user.Email
	pfpURL := user.PFPURL

	if username == "" || email == "" || pfpURL == "" {
		return fiber.StatusBadRequest, "Please attatch all fields of information (Username, Email, PFPURL)"
	}

	if method == "id" {
		request = fmt.Sprintf(
			`UPDATE "Users"
      SET "Username" = '%v', "Email" = '%v', "PFPURL" = '%v'
      WHERE "ID" = '%v';`,
			username, email, pfpURL, value)
	} else if method == "username" {
		request = fmt.Sprintf(
			`UPDATE "Users"
      SET "Username" = '%v', "Email" = '%v', "PFPURL" = '%v'
      WHERE "Username" = '%v';`,
			username, email, pfpURL, value)
	}

	_, err := database.DB.Exec(request)
	if err != nil {
		return fiber.StatusBadRequest, "Failed to execute database command, please recheck all fields of information. (Username and Email must be unique)"
	}

	return fiber.StatusOK, "Success"
}
