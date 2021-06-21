package utils

import "strings"

func Method(input string) string {
	if strings.ToLower(input) == "id" {
		return "ID"
	} else if strings.ToLower(input) == "title" {
		return "Title"
	} else if strings.ToLower(input) == "username" {
		return "Username"
	} else {
		return "ID"
	}
}
