package validation

import (
	"fmt"
	"strconv"
)

var isAnswerIncorrect bool

func IsAnswerIncorrect() bool {
	return isAnswerIncorrect
}

// Validation of numbers
func ValidateFloat(s string) (float64, bool) {
	number, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("'%s' is not a number\n", s)
		return 0, false
	} else if number <= 0 {
		fmt.Println("You have to input non-negative numbers!")
		return 0, false
	} else {
		return number, true
	}
}

// Validation of the entered as a triangle sides
func TriangleExistsValidator(sides []float64) bool {
	return sides[0]+sides[1] > sides[2] &&
		sides[0]+sides[2] > sides[1] &&
		sides[1]+sides[2] > sides[0]
}

func FinalTriangleValidator(splitSlice []string) ([]float64, []bool) {
	checker := make([]bool, 4)
	floats := make([]float64, 3)
	var isWrongInput bool

	if len(splitSlice[0]) <= 0 { // Name of the triangle can't be empty
		checker[0] = false
		fmt.Println("Name of triangle can't be empty")
	} else {
		checker[0] = true
	}

	for i, j := 1, 0; i < 4; i, j = i+1, j+1 { // validation of the entered numbers
		floats[j], checker[i] = ValidateFloat(splitSlice[i]) // If at least one value is unacceptable -
		if checker[i] == false {                             // return from validator and notification about the error
			isWrongInput = true
		}
	}

	if isWrongInput {
		return floats, checker
	}

	if !TriangleExistsValidator(floats) {
		for i, j := 1, 0; i < 4; i, j = i+1, j+1 {
			checker[i], floats[j] = false, 0
		}
		fmt.Println("Such a triangle cannot exist")
	}

	return floats, checker
}

// Validation of user's answer
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
