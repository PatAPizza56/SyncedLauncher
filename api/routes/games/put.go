package games

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber"

	database "../../database"
	structs "../../structs"
)

func Put(c *fiber.Ctx) error {
	game := new(structs.Game)

	err := c.BodyParser(game)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Failed to parse body"))
	}

	status, message := update(c.Query("method"), strings.ReplaceAll(c.Params("value"), "%20", " "), *game)

	c.Status(status)
	return c.Send([]byte(message))
}

func update(method string, value string, game structs.Game) (int, string) {
	var request string

	userID := game.UserID
	title := game.Title
	description := game.Description
	price := game.Price
	downloadURL := game.DownloadURL
	donationURL := game.DonationURL
	bannerURL := game.BannerURL

	if userID == 0 || title == "" || description == "" || price == "" || downloadURL == "" || donationURL == "" || bannerURL == "" {
		return fiber.StatusBadRequest, "Please attatch all fields of information (UserID, Title, Description, Price, DownloadURL, DonationURL, BannerURL)"
	}

	if method == "id" {
		request = fmt.Sprintf(
			`UPDATE "Games"
      SET "UserID" = '%v', "Title" = '%v', "Description" = '%v', "Price" = '%v', "DownloadURL" = '%v', "DonationURL" = '%v', "BannerURL" = '%v'
      WHERE "ID" = '%v';`,
			userID, title, description, price, downloadURL, donationURL, bannerURL, value)
	} else if method == "title" {
		request = fmt.Sprintf(
			`UPDATE "Games"
      SET "UserID" = '%v', "Title" = '%v', "Description" = '%v', "Price" = '%v', "DownloadURL" = '%v', "DonationURL" = '%v', "BannerURL" = '%v'
      WHERE "Title" = '%v';`,
			userID, title, description, price, downloadURL, donationURL, bannerURL, value)
	}

	_, err := database.DB.Exec(request)
	if err != nil {
		return fiber.StatusBadRequest, "Failed to execute database command, please recheck all fields of information. (Title and DownloadURL must be unique)"
	}

	return fiber.StatusOK, "Success"
}
