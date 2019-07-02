package main

import (
	"strings"
	"unicode"
)

// Removing all spaces from the string. Replace(), TrimSpace() can't provide needed functional
func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return -1
		}
		// else keep it in the string
		return r
	}, str)
}
