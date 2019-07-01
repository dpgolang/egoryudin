// Программа позволяет вывести все числа Фибоначчи, которые находятся в указанном диапазоне.
// Диапазон задаётся двумя аргументами. Числа выводятся через запятую по возрастанию.
package main

import (
	"Eighth/validation"
	"bufio"
	"fmt"
	"os"
)

func isEndBiggerThanBegin(fib []int64) bool {
	if fib[1] <= fib[0] {
		fmt.Println("End of range must be bigger than the beginning")
		return false
	}
	return true
}

func calcAndShowFiboSequence(fib []int64) {
	prev, next := int64(0), int64(1)
	for i := 0; ; i++ {
		temp := next
		next = prev + next
		prev = temp
		if next >= fib[0] && next <= fib[1] {
			fmt.Printf("%d, ", next)
		} else if next > fib[1] {
			break
		}
	}
}

func main() {
	var (
		scanner     = bufio.NewScanner(os.Stdin)
		input       string
		fib         = make([]int64, 2)
		isFibNumber = make(boolSlice, 2)
	)
	fmt.Println("Please input range limit.\n" +
		"First number: ")
	scanner.Scan()
	input = scanner.Text()
	if isFibNumber[0], fib[0] = validation.ValidateInt(input); !isFibNumber[0] {
		return
	}

	fmt.Println("Second number: ")
	scanner.Scan()
	input = scanner.Text()
	if isFibNumber[1], fib[1] = validation.ValidateInt(input); !isFibNumber[1] {
		return
	}

	if isFibNumber.contains(false) || !isEndBiggerThanBegin(fib) {
		return
	}

	calcAndShowFiboSequence(fib)
}
