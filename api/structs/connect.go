package structs

import (
	"strconv"

	"../database"
)

type Connect struct {
	ID        int
	UserID    int
	ConnectID string
}

func (connect *Connect) Post(id *int) (error, int) {
	fields := []string{"UserID", "ConnectID"}
	values := []string{strconv.Itoa(connect.UserID), connect.ConnectID}

	err, stat := database.Insert("Connect", fields, values, id)
	return err, stat
}

func (connect *Connect) Get(param string, value string) (error, int) {
	err, stat := database.Select("Connect", param, value, &connect.ID, &connect.UserID, &connect.ConnectID)
	return err, stat
}

func (connect *Connect) Put(param string, value string) (error, int) {
	fields := []string{"UserID", "ConnectID"}
	values := []string{strconv.Itoa(connect.UserID), connect.ConnectID}

	err, stat := database.Update("Connect", param, value, fields, values)
	return err, stat
}

func (connect *Connect) Delete(param string, value string) (error, int) {
	err, stat := database.Delete("Connect", param, value)
	return err, stat
}
