package games

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber"

	database "../../database"
	structs "../../structs"
)

func Get(c *fiber.Ctx) error {
	var user structs.User

	status, message := fetch(c.Query("method"), strings.ReplaceAll(c.Params("value"), "%20", " "), &user)
	c.Status(status)
	if message != "" {
		return c.Send([]byte(message))
	}

	return c.JSON(user)
}

func fetch(method string, value string, user *structs.User) (int, string) {
	var request string

	var ID int
	var Username string
	var Email string
	var Password string
	var PFPURL string

	if method == "id" {
		request = fmt.Sprintf(`SELECT * FROM "Users" WHERE "ID" = '%v'`, value)
	} else if method == "username" {
		request = fmt.Sprintf(`SELECT * FROM "Users" WHERE "Username" = '%v'`, value)
	}

	err := database.DB.QueryRow(request).Scan(&ID, &Username, &Email, &Password, &PFPURL)
	if err != nil {
		return fiber.StatusNotFound, "Failed to located user, please recheck all fields of information"
	}

	Password = ""
	*user = structs.User{ID: ID, Username: Username, Email: Email, PFPURL: PFPURL}
	return fiber.StatusOK, ""
}
