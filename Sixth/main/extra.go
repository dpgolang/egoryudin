package main

import (
	"strings"
	"unicode"
)

// Удаление пробелов в строках. Replace(), TrimSpace() не дают полного удаления всех пробелов
func stripSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// если символ пробел, тогда убрать его
			return -1
		}
		// иначе оставить символ в стркое
		return r
	}, str)
}
