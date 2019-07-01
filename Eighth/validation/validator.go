package validation

import (
	"fmt"
	"math"
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
		}

		return true, number
	} else {
		fmt.Print("To start a program you have to input length and width of the envelopes!")
		return false, 0
	}
}

// A utility function that returns true if x is perfect square
func isPerfectSquare(x float64) bool {
	s := math.Sqrt(x)
	return (s * s) == x
}

// Returns true if n is a Fibonacci Number, else false
func ValidateFibNumber(s string) (bool, int64) {
	if isNumberCorrect, fib := ValidateInt(s); isNumberCorrect {
		// fib is Fibonacci if one of 5*n*n + 4 or 5*n*n - 4 or both
		// is a perfect square
		if isPerfectSquare(5*float64(fib)*float64(fib)+4) ||
			isPerfectSquare(5*float64(fib)*float64(fib)-4) {
			return true, fib
		} else {
			fmt.Println("Input value cannot be a Fibonacci number")
			return false, 0
		}
	}
	return false, 0
}
