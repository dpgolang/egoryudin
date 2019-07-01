package main

import (
	"Third/validation"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var stdin = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(os.Stdin)

// Ф-ия добавления треугольников
func AddTriangle() []Triangle {
	incorrectChecker := make(boolSlice, 4) // Слайс, элементы которого показывают, все ли значения введены корректно
	splitValues := make([]string, 4)       // Слайс в котором хранятся имя и три стороны треугольника
	tempFloats := make([]float64, 3)       // Слайс для временного хранения значения введённых сторон треугольника
	var triangleSlice []Triangle           // Слайс, в котором хранятся все треугольники
	var inputtedValue string               // Введённая строка с консоли
	var isAnyIncorrectValue = true         // Переменная, которая отвечает за то, было ли введено хотя бы одно недопустимое значение

	// Цикл добавления новых треугольников. Длится до тех пор, пока пользователь не напишет 'n/no'
	for {
		// Цикл ввода значений треугольника. Длится до тех пор, пока все значения не будут допустимыми
		for isAnyIncorrectValue {
			fmt.Println("\nInput name and lengths of the three sides of triangle")
			scanner.Scan()
			inputtedValue = scanner.Text()

			// Разбиение строки с введёнными значениями на слайс
			splitValues = strings.Split(inputtedValue, ",")
			if len(splitValues) != 4 {
				fmt.Println("You have to input 4 values!")
				continue
			}
			// Удаление всех пробелов в значениях слайса
			for i := 0; i < 4; i++ {
				splitValues[i] = stripSpaces(splitValues[i])
			}
			incorrectChecker, tempFloats = validation.FinalTriangleValidator(splitValues) // Полная проверка введённых значений
			isAnyIncorrectValue = incorrectChecker.contains()                             // Проверка было ли введено хотя бы одно недопустимое значение
			incorrectChecker = incorrectChecker[:0]                                       // Очистка чекера введённых значений
		}
		triangleSlice = append(triangleSlice, Triangle{splitValues[0], tempFloats[0], tempFloats[1], tempFloats[2], 0}) // Добавление нового треугольинка

		if IfContinueInput() { // Если пользователь хочет продолжить, он вводит ещё один треугольник
			isAnyIncorrectValue = true
			continue
		} else { // Если нет - высчитывается площадь всех треугольников
			for i := 0; i < len(triangleSlice); i++ {
				triangleSlice[i].HeronsFormula()
			}
			fmt.Println("Triangle/triangles has/have been successfully added!\n")
			return triangleSlice // Возвращение в меню
		}
	}
}

// Сортироровка и вывод значений треугольников
func SortAndShowTriangles(triangles []Triangle) {
	if len(triangles) == 0 {
		fmt.Println("No triangles have been added\n")
		return
	}

	for i := 0; i < len(triangles); i++ {
		for j := i + 1; j < len(triangles); j++ {
			if triangles[i].s < triangles[j].s {
				temp := triangles[i]
				triangles[i] = triangles[j]
				triangles[j] = temp
			}
		}
	}

	fmt.Println("\n============= Triangles list: ===============")
	for i := 0; i < len(triangles); i++ {
		fmt.Printf("[Triangle %s]: %f cm\n", triangles[i].name, triangles[i].s)
	}
	fmt.Println()
}

// Главное меню
func main() {
	var FilledTriangleSlice []Triangle
	var firstOption int
	for {
		fmt.Println("Please select an option\n" +
			"1. Add triangle\n" +
			"2. Sort and show list of triangles\n" +
			"3. Exit program")
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
			FilledTriangleSlice = AddTriangle()
		case 2:
			SortAndShowTriangles(FilledTriangleSlice)
		default:
			fmt.Println("Unknown option.")
		}
	}
}
