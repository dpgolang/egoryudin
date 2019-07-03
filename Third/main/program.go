package main

import (
	"Temp/Third/validation"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var stdin = bufio.NewReader(os.Stdin)
var scanner = bufio.NewScanner(os.Stdin)

// Adding new triangles
func AddTriangle() []Triangle {
	var (
		incorrectChecker    = make(boolSlice, 4) // Slice, elements of which show whether all values are entered correctly
		splitValues         = make([]string, 4)  // Slice which stores the name and the three sides of the triangle
		tempFloats          = make([]float64, 3) // Slice for temporary storage of the entered values of the sides of the triangle
		triangleSlice       []Triangle           // Slice in which all triangles are sorted
		inputValue          string               // Input string from console
		isAnyIncorrectValue = true               // A variable that is responsible for whether at least one invalid value was entered
	)

	// The cycle of adding new triangles. Lasts until the user writes 'n/no'
	for {
		// The cycle for entering the values of a triangle. Lasts until all values are valid
		for isAnyIncorrectValue {
			fmt.Println("\nInput name and lengths of the three sides of triangle")
			scanner.Scan()
			inputValue = scanner.Text()

			// Splitting a string with entered values into a slice
			splitValues = strings.Split(inputValue, ",")
			if len(splitValues) != 4 {
				fmt.Println("You have to input 4 values!")
				continue
			}
			// Removing all spaces in slice elements
			for i := 0; i < 4; i++ {
				splitValues[i] = stripSpaces(splitValues[i])
			}
			tempFloats, incorrectChecker = validation.FinalTriangleValidator(splitValues) // Full validation of entered values
			isAnyIncorrectValue = incorrectChecker.contains()                             // Check if at least one invalid value was entered
			incorrectChecker = incorrectChecker[:0]                                       // Clearing the checker
		}
		triangleSlice = append(triangleSlice, Triangle{splitValues[0], tempFloats[0], tempFloats[1], tempFloats[2], 0}) // Adding new triangle

		if IfContinueInput() { // If the user would like to continue, he adds another triangle
			isAnyIncorrectValue = true
			continue
		} else { // Else - the area of all triangles calculates
			for i := 0; i < len(triangleSlice); i++ {
				triangleSlice[i].HeronsFormula()
			}
			fmt.Println("Triangle/triangles has/have been successfully added!\n")
			return triangleSlice // Return to the menu
		}
	}
}

// Sorting and output triangles
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

// Main menu
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
