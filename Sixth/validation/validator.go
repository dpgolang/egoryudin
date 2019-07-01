package validation

import (
	"fmt"
	"strconv"
)

func ValidateInt(s string) (bool, int) {
	if len(s) == 6 {
		number, err := strconv.Atoi(s)
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
		fmt.Printf("Ticket number '%s' must be a length of 6.\n", s)
		return false, 0
	}
}

func ValidateCity(s string) (bool, bool) { // Первое возвращаемое значение показывает, введенна ли допустимая строка, второе - алгоритм "Moscow" или "Petersburg"
	if s == "Moscow" {
		return true, true
	} else if s == "Petersburg" {
		return true, false
	}
	return false, false
}
