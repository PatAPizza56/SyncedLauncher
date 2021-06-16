package main

import (
	database "./database"
	rest "./rest"
)

func main() {
	database.Open()
	rest.Start()
}
