package validation

import (
	"fmt"
	"strconv"
)

func ValidateInput(s string) (bool, uint) {
	if len(s) > 0 {
		number, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("'%s' is not a number\n", s)
			return false, 0
		} else if number <= 0 {
			fmt.Println("You have to input non-negative numbers!")
			return false, 0
		} else {
			return true, uint(number)
		}
	} else {
		fmt.Println("To start a program you have to input length and width of the chessboard!")
		return false, 0
	}
}
