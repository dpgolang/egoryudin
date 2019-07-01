package main

import (
	"Fifth/validation"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func wordChooser(number int64) string {
	var (
		numLength = len(strconv.FormatInt(number, 10))
	)

	if numLength == 1 {
		return getDigitNumString(number)
	} else if numLength == 2 {
		return getTenOrTeenNumString(number)
	} else if numLength == 3 {
		return getHundredNumString(number)
	} else if numLength > 3 && numLength < 7 {
		return getThousandNumString(number)
	} else if numLength >= 7 && numLength < 10 {
		return getMillionNumString(number)
	} else if numLength >= 10 && numLength < 13 {
		return getBillionNumString(number)
	} else {
		fmt.Println("Error.")
	}
	return ""
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := scanner.Text()
		if isInputCorrect, number := validation.ValidateInt(input); isInputCorrect {
			fmt.Println(wordChooser(number))
		}
	}
}
