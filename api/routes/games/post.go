package games

import (
	"fmt"

	"github.com/gofiber/fiber"

	database "../../database"
	structs "../../structs"
)

func Post(c *fiber.Ctx) error {
	game := new(structs.Game)

	err := c.BodyParser(game)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.Send([]byte("Failed to parse body"))
	}

	status, message := insert(*game)

	c.Status(status)
	return c.Send([]byte(message))
}

func insert(game structs.Game) (int, string) {
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

	request := fmt.Sprintf(
		`INSERT INTO "Games" ("UserID", "Title", "Description", "Price", "DownloadURL", "DonationURL", "BannerURL")
    VALUES ('%v', '%v', '%v', '%v', '%v', '%v', '%v');`,
		userID, title, description, price, downloadURL, donationURL, bannerURL)

	_, err := database.DB.Exec(request)
	if err != nil {
		return fiber.StatusBadRequest, "Failed to execute database command, please recheck all fields of information. (Title and DownloadURL must be unique)"
	}

	return fiber.StatusOK, "Success"
}
