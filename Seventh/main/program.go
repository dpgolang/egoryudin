// Программа выводит ряд натуральных чисел через запятую, квадрат которых меньше заданного n..
package main

import (
	"Seventh/validation"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		scanner        = bufio.NewScanner(os.Stdin) // input from console
		input          string                       // input string
		n              int64                        // n parameter
		isInputCorrect bool                         // checking if input string can be parsed
	)

	fmt.Println("Input parameter (n)")
	scanner.Scan()
	input = scanner.Text()

	if isInputCorrect, n = validation.ValidateInt(input); isInputCorrect {
		for i := int64(1); i*i < n; i++ {
			fmt.Print(i, ", ")
		}
	}
}
