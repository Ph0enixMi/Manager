package utils

import "unicode/utf8"

func GetSpaces(text string) string {
	var spaces string
	lenght := utf8.RuneCountInString(text)
	for range lenght {
		spaces += " "
	}
	return spaces
}
