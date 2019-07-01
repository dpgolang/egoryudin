package validation

import (
	"fmt"
	"strconv"
)

// Parsing input string
func ValidateInt(s string) (bool, int64) {
	if len(s) > 0 {
		number, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			// Parsing if parsed int64 is out of range
			if numError, ok := err.(*strconv.NumError); ok {
				if numError.Err == strconv.ErrRange {
					fmt.Println("Detected", numError.Num, "as a", strconv.ErrRange)
					return false, 0
				}
			}

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
