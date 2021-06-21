package structs

import (
	"strconv"

	"../database"
)

type Token struct {
	ID     int
	UserID int
	Value  string
}

func (token *Token) Post(id *int) (error, int) {
	fields := []string{"UserID", "Value"}
	values := []string{strconv.Itoa(token.UserID), token.Value}

	err, stat := database.Insert("Tokens", fields, values, id)
	return err, stat
}

func (token *Token) Get(param string, value string) (error, int) {
	err, stat := database.Select("Tokens", param, value, &token.ID, &token.UserID, &token.Value)
	return err, stat
}

func (token *Token) Put(param string, value string) (error, int) {
	fields := []string{"UserID", "Value"}
	values := []string{strconv.Itoa(token.UserID), token.Value}

	err, stat := database.Update("Tokens", param, value, fields, values)
	return err, stat
}

func (token *Token) Delete(param string, value string) (error, int) {
	err, stat := database.Delete("Tokens", param, value)
	return err, stat
}
