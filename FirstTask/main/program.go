package main

import (
	"FirstTask/validation"
	"bufio"
	"fmt"
	"os"
)

var chessBoard Board
var stdin = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(os.Stdin)

//ф-ия, в которой задаём размеры шахматной доски, предварительно проверяя введённые данные
func SetChessboard() {

	for {
		var inputLength, inputWidth string
		var isNumberChecked bool
		fmt.Println("\nPlease, enter the length of chessboard")
		scanner.Scan()
		inputLength = scanner.Text()
		if isNumberChecked, chessBoard.Length = validation.ValidateInput(inputLength); !isNumberChecked {
			continue
		}
		fmt.Println("Enter the width of chessboard")
		scanner.Scan()
		inputWidth = scanner.Text()
		if isNumberChecked, chessBoard.Width = validation.ValidateInput(inputWidth); !isNumberChecked {
			continue
		}
		break
	}
	fmt.Printf("Length of chessboard = %d\n"+
		"Width of chessboard = %d\n\n", chessBoard.Length, chessBoard.Width)
}

//ф-ия вывода шахматной доски с заданными размерами
func DrawChessboard() {
	var s string
	if chessBoard.Width == 0 || chessBoard.Length == 0 {
		s = "Firstly, you have to input sizes of the chessboard.\n" +
			"Please, select 1 option in main menu.\n\n"
	} else {
		for i := 1; i < int(chessBoard.Width)+1; i++ {
			if i%2 == 0 {
				s += " "
			}
			for j := 1; j < int(chessBoard.Length)+1; j++ {
				s += "* "
			}
			s += "\n"
		}
	}
	fmt.Print(s)
}

//ф-ия меню
func main() {
	for true {
		var option int
		fmt.Println("Welcome to my chessboard!\n" +
			"Please select an option\n" +
			"1. Set chessboard sizes\n" +
			"2. Draw chessboard in the console\n" +
			"3. Exit program")
		for {
			_, err := fmt.Fscan(stdin, &option)
			if err == nil {
				break
			}
			stdin.ReadString('\n')
			fmt.Println("Sorry, invalid input. Please enter correct number of option: ")
		}

		switch option {
		case 1:
			SetChessboard()
		case 2:
			DrawChessboard()
		case 3:
			os.Exit(3)
		default:
			fmt.Println("Unknown option.")
		}
	}
}
