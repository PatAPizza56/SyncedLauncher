package structs

import (
	"strconv"

	"../database"
)

type Game struct {
	ID          int
	UserID      int
	Title       string
	Description string
	Price       string
	DownloadURL string
	BannerURL   string
}

func (game *Game) Post(id *int) (error, int) {
	fields := []string{"UserID", "Title", "Description", "Price", "DownloadURL", "BannerURL"}
	values := []string{strconv.Itoa(game.UserID), game.Title, game.Description, game.Price, game.DownloadURL, game.BannerURL}

	err, stat := database.Insert("Games", fields, values, id)
	return err, stat
}

func (game *Game) Get(param string, value string) (error, int) {
	err, stat := database.Select("Games", param, value, &game.ID, &game.UserID, &game.Title, &game.Description, &game.Price, &game.DownloadURL, &game.BannerURL)
	return err, stat
}

func (game *Game) Put(param string, value string) (error, int) {
	fields := []string{"UserID", "Title", "Description", "Price", "DownloadURL", "BannerURL"}
	values := []string{strconv.Itoa(game.UserID), game.Title, game.Description, game.Price, game.DownloadURL, game.BannerURL}

	err, stat := database.Update("Games", param, value, fields, values)
	return err, stat
}

func (game *Game) Delete(param string, value string) (error, int) {
	err, stat := database.Delete("Games", param, value)
	return err, stat
}
