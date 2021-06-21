package utils

import "strings"

func Value(input string) string {
	return strings.ReplaceAll(input, "%20", " ")
}
