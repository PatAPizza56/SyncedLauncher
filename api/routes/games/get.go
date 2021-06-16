package games

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber"

	database "../../database"
	structs "../../structs"
)

func Get(c *fiber.Ctx) error {
	var game structs.Game

	status, message := fetch(c.Query("method"), strings.ReplaceAll(c.Params("value"), "%20", " "), &game)

	c.Status(status)
	if message != "" {
		return c.Send([]byte(message))
	}

	return c.JSON(game)
}

func fetch(method string, value string, game *structs.Game) (int, string) {
	var request string

	var ID int
	var UserID int
	var Title string
	var Description string
	var Price string
	var DownloadURL string
	var DonationURL string
	var BannerURL string

	if method == "id" {
		request = fmt.Sprintf(`SELECT * FROM "Games" WHERE "ID" = '%v'`, value)
	} else if method == "title" {
		request = fmt.Sprintf(`SELECT * FROM "Games" WHERE "Title" = '%v'`, value)
	}

	err := database.DB.QueryRow(request).Scan(&ID, &UserID, &Title, &Description, &Price, &DownloadURL, &DonationURL, &BannerURL)
	if err != nil {
		return fiber.StatusNotFound, "Failed to located game, please recheck all fields of information"
	}

	*game = structs.Game{ID: ID, UserID: UserID, Title: Title, Description: Description, Price: Price, DownloadURL: DownloadURL, DonationURL: DonationURL, BannerURL: BannerURL}
	return fiber.StatusOK, ""
}
