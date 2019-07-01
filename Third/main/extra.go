package main

import (
	"Third/validation"
	"fmt"
	"strings"
	"unicode"
)

// Кастомный тип, имеющий одну ф-ию. Слайс типа bool проверяется на наличие хотя бы одного значения true.
type boolSlice []bool

func (b boolSlice) contains() bool {
	for _, a := range b {
		if !a {
			return true
		}
	}
	return false
}

// Поиск пробелов в строке и их удаление
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

// Ф-ия для продолжения/выхода из ф-ий
func IfContinueInput() bool {
	var UsersAnswer string
	var ToContinue bool
	for {
		fmt.Println("\nWould you like to continue? (y,yes/n,no)")

		scanner.Scan()
		UsersAnswer = scanner.Text()
		strings.ToLower(UsersAnswer)
		ToContinue = validation.ValidateAnswer(UsersAnswer)

		if validation.IsAnswerIncorrect() {
			fmt.Println("Incorrect input. Try again.\n")
			continue
		}
		if ToContinue {
			return true
		} else {
			fmt.Println()
			return false
		}
	}
}
