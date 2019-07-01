package main

import (
	"Temp/Second/validation"
	"fmt"
)

type Envelope struct {
	length, width float64
	num           int
}

// Вывод сторон конверта
func (e *Envelope) ShowInputtedSizes() {
	fmt.Printf("Length of %d envelope is %f\n", e.num, e.length)
	fmt.Printf("Width of %d envelope is %f\n", e.num, e.width)
	fmt.Println()
}

// Первый способ заполнения сторон конвертов
func (e *Envelope) SetEnvelopesSidesFirstMode() {
	var IsNumberChecked bool
	var InputtedValue string

	for i := 1; i < 3; i++ {
	InputSizes:
		// Аномниная ф-ия, отвечающая за заполнение длины или ширины
		LengthOrWidthFunc := func(i int) bool { return i&1 == 0 }
		if !LengthOrWidthFunc(i) {
			fmt.Printf("\n\nPlease, input length of %d envelope\n", e.num)
			e.length = 0
		} else {
			fmt.Printf("\nPlease, input width of %d envelope\n", e.num)
			e.width = 0
		}
		scanner.Scan()
		InputtedValue = scanner.Text()

		if !LengthOrWidthFunc(i) {
			if IsNumberChecked, e.length = validation.ValidateFloat(InputtedValue); !IsNumberChecked {
				goto InputSizes
			}
		} else {
			if IsNumberChecked, e.width = validation.ValidateFloat(InputtedValue); !IsNumberChecked {
				goto InputSizes
			}
		}

		// Если i == 2, это последняя итерация в цикле
		if i == 2 {
			fmt.Println("Sizes have been successfully saved!\n")
			return
		}

		if IfContinueInput() {
			continue
		} else {
			fmt.Println()
			return
		}
	}
}

// Второй способ заполнения сторон конвертов
func SetEnvelopesSidesSecondMode(envelopes []Envelope) {
	var InputtedValue string
	var IsNumberChecked bool
InputSizes:
	for i := 0; i < len(envelopes); i++ {
		for !IsNumberChecked {
			fmt.Printf("\nPlease, input length of %d envelope\n", envelopes[i].num)
			envelopes[i].length = 0
			scanner.Scan()
			InputtedValue = scanner.Text()

			IsNumberChecked, envelopes[i].length = validation.ValidateFloat(InputtedValue)
		}

		IsNumberChecked = false

		for !IsNumberChecked {
			fmt.Printf("\nPlease, input width of %d envelope\n", envelopes[i].num)
			envelopes[i].width = 0
			scanner.Scan()
			InputtedValue = scanner.Text()

			IsNumberChecked, envelopes[i].width = validation.ValidateFloat(InputtedValue)
		}

		IsNumberChecked = false
	}
	CompareEnvelopes(envelopes[0], envelopes[1])

	if IfContinueInput() {
		goto InputSizes
	}
	return
}
