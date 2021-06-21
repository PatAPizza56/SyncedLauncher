package structs

import (
	"errors"
	"strconv"

	"../database"
	"github.com/gofiber/fiber"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	PFPURL   string
}

func (user *User) Post(id *int) (error, int) {
	fields := []string{"Username", "Email", "Password", "PFPURL"}
	values := []string{user.Username, user.Email, user.Password, user.PFPURL}

	err, stat := database.Insert("Users", fields, values, id)
	return err, stat
}

func (user *User) Get(param string, value string) (error, int) {
	err, stat := database.Select("Users", param, value, &user.ID, &user.Username, &user.Email, &user.Password, &user.PFPURL)
	return err, stat
}

func (user *User) GetByToken(value string) (error, int) {
	var token Token

	err, _ := token.Get("Value", value)
	if err != nil {
		return errors.New("Failed to validate token"), fiber.StatusBadRequest
	}

	err, _ = database.Select("Users", "ID", strconv.Itoa(token.UserID), &user.ID, &user.Username, &user.Email, &user.Password, &user.PFPURL)
	if err != nil {
		return errors.New("Failed to validate token"), fiber.StatusBadRequest
	}

	return nil, 0
}

func (user *User) Put(param string, value string) (error, int) {
	fields := []string{"Username", "Email", "Password", "PFPURL"}
	values := []string{user.Username, user.Email, user.Password, user.PFPURL}

	err, stat := database.Update("Users", param, value, fields, values)
	return err, stat
}

func (user *User) Delete(param string, value string) (error, int) {
	err, stat := database.Delete("Users", param, value)
	return err, stat
}
