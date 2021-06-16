package games

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber"

	database "../../database"
)

func Delete(c *fiber.Ctx) error {
	status, message := remove(c.Query("method"), strings.ReplaceAll(c.Params("value"), "%20", " "))

	c.Status(status)
	return c.Send([]byte(message))
}

func remove(method string, value string) (int, string) {
	var request string

	if method == "id" {
		request = fmt.Sprintf(`DELETE FROM "Users" WHERE "ID" = '%v'`, value)
	} else if method == "username" {
		request = fmt.Sprintf(`DELETE FROM "Users" WHERE "Username" = '%v'`, value)
	}

	_, err := database.DB.Exec(request)
	if err != nil {
		return fiber.StatusBadRequest, "Failed to located user, please recheck all fields of information"
	}

	return fiber.StatusOK, "Success"
}
