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
func ValidateFloat(s string) (bool, float64){
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
}

// Проверка, может ли треугольник с введёнными сторонами существовать
func TriangleExistsValidator(sides []float64) bool{
	return sides[0] + sides[1] > sides[2] &&
		sides[0] + sides[2] > sides[1] &&
		sides[1] + sides[2] > sides[0]
}

// Общая проверка треугольника
func FinalTriangleValidator(triangleSlice []string) ([]bool, []float64) {
	checker := make([]bool, 4)
	floats := make([]float64, 3)
	var isWrongInput bool

	if len(triangleSlice[0]) <= 0 {	// Имя треугольника не может быть пустым
		checker[0] = false
		fmt.Println("Name of triangle can't be empty")
	} else {
		checker[0] = true
	}

	for i, j := 1, 0; i < 4; i, j = i + 1, j + 1 {			// Проверка введённых значений сторон треугольника
		checker[i], floats[j] = ValidateFloat(triangleSlice[i]) // Если хотя бы одно значение недопустимо -
		if checker[i] == false {				// выход из валидатора и оповещение об ошибке
			isWrongInput = true
		}
	}

	if isWrongInput {
		return checker, floats
	}

	if !TriangleExistsValidator(floats) {
		for i, j := 1, 0; i < 4; i, j = i + 1, j + 1 {
			checker[i], floats[j] = false, 0
		}
		fmt.Println("Such a triangle cannot exist")
	}

	return checker, floats
}

//Проверка корректности ответа пользователя
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