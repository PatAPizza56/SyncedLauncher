package main

import (
	"./database"
	"./rest"
)

func main() {
	database.Open()
	rest.Start()
}
