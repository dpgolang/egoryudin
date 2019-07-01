package validation

import (
	"fmt"
	"strconv"
)

var isAnswerIncorrect bool

func IsAnswerIncorrect() bool {
	return isAnswerIncorrect
}

// Проверка корректности введённых значений
func ValidateFloat(s string) (bool, float64) {
	if len(s) > 0 {
		number, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Printf("'%s' is not a number\n", s)
			return false, 0
		} else if number <= 0 {
			fmt.Println("You have to input non-negative numbers!")
			return false, 0
		} else {
			return true, number
		}
	} else {
		fmt.Print("To start a program you have to input length and width of the envelopes!")
		return false, 0
	}
}

// Проверка корректности ответа пользователя
func ValidateAnswer(answer string) bool {
	isAnswerIncorrect = false
	if answer == "y" || answer == "yes" {
		return true
	} else if answer == "n" || answer == "no" {
		return false
	} else {
		isAnswerIncorrect = true
		return false
	}
}
