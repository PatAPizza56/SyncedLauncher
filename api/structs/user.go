package structs

import (
	"errors"
	"strconv"

	"../database"
	"github.com/gofiber/fiber"
)

type User struct {
	ID       int
	FullName string
	Username string
	Email    string
	Password string
	PfpURL   string
}

func (user *User) Post(id *int) (error, int) {
	fields := []string{"FullName", "Username", "Email", "Password", "PfpURL"}
	values := []string{user.FullName, user.Username, user.Email, user.Password, user.PfpURL}

	err, stat := database.Insert("Users", fields, values, id)
	return err, stat
}

func (user *User) Get(param string, value string) (error, int) {
	err, stat := database.Select("Users", param, value, &user.ID, &user.FullName, &user.Username, &user.Email, &user.Password, &user.PfpURL)
	return err, stat
}

func (user *User) GetByToken(value string) (error, int) {
	var token Token

	err, _ := token.Get("Value", value)
	if err != nil {
		return errors.New("Failed to validate token"), fiber.StatusBadRequest
	}

	err, _ = database.Select("Users", "ID", strconv.Itoa(token.UserID), &user.ID, &user.FullName, &user.Username, &user.Email, &user.Password, &user.PfpURL)
	if err != nil {
		return errors.New("Failed to validate token"), fiber.StatusBadRequest
	}

	return nil, 0
}

func (user *User) Put(param string, value string) (error, int) {
	fields := []string{"FullName", "Username", "Email", "Password", "PfpURL"}
	values := []string{user.FullName, user.Username, user.Email, user.Password, user.PfpURL}

	err, stat := database.Update("Users", param, value, fields, values)
	return err, stat
}

func (user *User) Delete(param string, value string) (error, int) {
	err, stat := database.Delete("Users", param, value)
	return err, stat
}
