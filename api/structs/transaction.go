package structs

import (
	"strconv"

	"../database"
)

type User_Game struct {
	ID     int
	UserID int
	GameID int
}

func (userGame *User_Game) Post(id *int) (error, int) {
	fields := []string{"UserID", "GameID"}
	values := []string{strconv.Itoa(userGame.UserID), strconv.Itoa(userGame.GameID)}

	err, stat := database.Insert("User_Game", fields, values, id)
	return err, stat
}

func (userGame *User_Game) Get(param string, value string) (error, int) {
	err, stat := database.Select("User_Game", param, value, &userGame.ID, &userGame.UserID, &userGame.GameID)
	return err, stat
}

func (userGame *User_Game) Put(param string, value string) (error, int) {
	fields := []string{"UserID", "GameID"}
	values := []string{strconv.Itoa(userGame.UserID), strconv.Itoa(userGame.GameID)}

	err, stat := database.Update("User_Game", param, value, fields, values)
	return err, stat
}

func (userGame *User_Game) Delete(param string, value string) (error, int) {
	err, stat := database.Delete("User_Game", param, value)
	return err, stat
}
