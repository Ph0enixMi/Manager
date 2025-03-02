package utils

import "unicode/utf8"

func GetSpacesBar(text string) string {
	var spaces string
	lenght := utf8.RuneCountInString(text)
	for range lenght {
		spaces += " "
	}
	return spaces
}

func GetMaxLn(list []string) int {
	mx := len(list[0])
	for _, elem := range list {
		if len(elem) > mx {
			mx = len(elem)
		}
	}
	return mx
}

func GetSpaces(text string, mx int) string {
	var spaces string
	if mx > len(text) {
		for range mx - len(text) {
			spaces += " "
		}
	}
	return spaces
}
