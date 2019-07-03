package main

import (
	"Temp/Third/validation"
	"fmt"
	"strings"
	"unicode"
)

// A custom type that has one function. A bool slice is checked for at least one true value.
type boolSlice []bool

func (b boolSlice) contains() bool {
	for _, a := range b {
		if !a {
			return true
		}
	}
	return false
}

// Searching for the spaces in string and removing them
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

// Continue/exit from function
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
