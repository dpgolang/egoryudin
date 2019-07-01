package main

import (
	"SecondTask/validation"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var stdin = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(os.Stdin)

// Проверка на то, можно ли поместить один конверт в другой и наоборот
func CompareEnvelopes(firstEnvelope, secondEnvelope Envelope) {
	if firstEnvelope.length == 0 || firstEnvelope.width == 0 || secondEnvelope.length == 0 || secondEnvelope.width == 0 {
		fmt.Println("Sides haven't been determined correctly")
		return
	}
	var isPut bool
	if CarversCondition(firstEnvelope, secondEnvelope) {
		fmt.Println("First envelope can be put into the second one\n")
		isPut = true
	}
	if CarversCondition(secondEnvelope, firstEnvelope) {
		fmt.Println("Second envelope can be put into the first one\n")
		isPut = true
	}
	if !isPut {
		fmt.Println("Envelopes can't be put into each other\n")
	}
}

// Формула проверки занесения конвертов один в другой
func CarversCondition(firstEnvelope, secondEnvelope Envelope) bool {
	firstBigSide := math.Max(firstEnvelope.length, firstEnvelope.width)
	firstSmallSide := math.Min(firstEnvelope.length, firstEnvelope.width)
	secondBigSide := math.Max(secondEnvelope.length, secondEnvelope.width)
	secondSmallSide := math.Min(secondEnvelope.length, secondEnvelope.width)

	if (firstBigSide <= secondBigSide && firstSmallSide <= secondSmallSide) ||
		(firstBigSide > secondBigSide && firstSmallSide <= secondSmallSide && math.Pow((secondBigSide+secondSmallSide)/(firstBigSide+firstSmallSide), 2)+
			math.Pow((secondBigSide-secondSmallSide)/(firstBigSide-firstSmallSide), 2) >= 2) {
		return true
	} else {
		return false
	}
}

// Первый режим работы программы.
// Пользователю будет разрешено выбрать конверт для заполнения (1/2).
// После каждого введённого значения пользователю будет предложено продолжить ввод.
// После занесения данных, в главном меню будет возможно попробовать вложить один конверт в другой.
func FirstMode(firstEnvelope *Envelope, secondEnvelope *Envelope) {
	firstEnvelope.length, firstEnvelope.width, secondEnvelope.length, secondEnvelope.width = 0, 0, 0, 0
	var firstOption int
	var secondOption int
	for {
		fmt.Println("Please select an option\n" +
			"1. Set envelopes sides\n" +
			"2. Show inputted sides\n" +
			"3. Try to put one envelope into another one\n" +
			"4. Change mode\n" +
			"5. Exit program")
		for {
			_, err := fmt.Fscan(stdin, &firstOption)
			if err == nil {
				break
			}
			stdin.ReadString('\n')
			fmt.Println("Sorry, invalid input. Please enter correct number of option: ")
		}

		switch firstOption {
		case 1:
			fmt.Println("\nChoose envelope, sizes of which you would like to fill (1/2).")
			for {
				_, err := fmt.Fscan(stdin, &secondOption)
				if err == nil {
					break
				}
				stdin.ReadString('\n')
				fmt.Println("Sorry, invalid input. Please enter correct number of option: ")
			}
			switch secondOption {
			case 1:
				firstEnvelope.SetEnvelopesSidesFirstMode()
			case 2:
				secondEnvelope.SetEnvelopesSidesFirstMode()
			default:
				fmt.Println("Unknown command.\n")
			}

		case 2:
			firstEnvelope.ShowInputtedSizes()
			secondEnvelope.ShowInputtedSizes()
		case 3:
			CompareEnvelopes(*firstEnvelope, *secondEnvelope)
		case 4:
			fmt.Println()
			return
		case 5:
			os.Exit(5)
		default:
			fmt.Println("Unknown command.\n")
		}
	}
}

// *Второй режим - оригинальное задание*
// Пользователю нельзя выбрать, какой конверт заполнять, он сразу  заполняет оба.
// После заполнения программа мгновенно покажет, можно ли положить один из конвертов в другой.
// Затем у пользователя спрашивается, хочет ли он продолжить.
// Согласие означает заполнение сторон конвертов с самого начала.
func SecondMode(firstEnvelope *Envelope, secondEnvelope *Envelope) {
	firstEnvelope.length, firstEnvelope.width, secondEnvelope.length, secondEnvelope.width = 0, 0, 0, 0
	var envelopes = make([]Envelope, 2)
	envelopes[0] = *firstEnvelope
	envelopes[1] = *secondEnvelope
	var firstOption int
	for {
		fmt.Println("Please select an option\n" +
			"1. Set envelopes sides and try to put one envelope into another one\n" +
			"2. Show inputted sides\n" +
			"3. Change mode\n" +
			"4. Exit program")

		for {
			_, err := fmt.Fscan(stdin, &firstOption)
			if err == nil {
				break
			}
			stdin.ReadString('\n')
			fmt.Println("Sorry, invalid input. Please enter correct number of option: ")
		}
		switch firstOption {
		case 1:
			SetEnvelopesSidesSecondMode(envelopes)
		case 2:
			envelopes[0].ShowInputtedSizes()
			envelopes[1].ShowInputtedSizes()
		case 3:
			fmt.Println()
			return
		case 4:
			os.Exit(4)
		}
	}
}

// Проверка, хочет ли пользователь продолжить работу
func IfContinueInput() bool {
	var UsersAnswer string
	var ToContinue bool
	for {
		fmt.Println("Would you like to continue? (y,yes/n,no)")

		scanner.Scan()
		UsersAnswer = scanner.Text()
		strings.ToLower(UsersAnswer)
		ToContinue = validation.ValidateAnswer(UsersAnswer)

		if validation.IsAnswerIncorrect() {
			fmt.Println("Incorrect input. Try again.\n")
			continue
		}
		if ToContinue {
			return true
		} else {
			fmt.Println()
			return false
		}
	}
}

// Главное меню, выбор режима работы программы
func main() {
	var firstEnvelope = Envelope{0, 0, 1}
	var secondEnvelope = Envelope{0, 0, 2}

	fmt.Println("Welcome to envelope task!")

	for {
		var chooseMode int

		fmt.Println("Please, choose program mode:\n" +
			"1. You will be allowed to choose which envelope to fill.\n" +
			"After each input you will be asked if you would like to continue.\n" +
			"Then in main menu you will be allowed to choose an option to see,\n" +
			"if one of the envelopes can be but into another one\n\n" +

			"2. You will not be allowed to choose which envelope to fill, you will\n" +
			"fill both of them right away.\n" +
			"After filling, program will instantly show if one of the envelopes\n" +
			"can be put into another one.\n" +
			"After that, you'll be asked if you would like to continue\n" +
			"'Yes' means filling sides of envelopes from the very beginning.\n\n" +
			"*Second mode is original task.")

		for {
			_, err := fmt.Fscan(stdin, &chooseMode)
			if err == nil {
				break
			}
			stdin.ReadString('\n')
			fmt.Println("Sorry, invalid input. Please enter correct number of mode: ")
		}
		fmt.Println()

		switch chooseMode {
		case 1:
			FirstMode(&firstEnvelope, &secondEnvelope)
		case 2:
			SecondMode(&firstEnvelope, &secondEnvelope)
		default:
			fmt.Println("Unknown option")
		}
	}
}
