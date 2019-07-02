package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var stdin = bufio.NewReader(os.Stdin)

// Setting .txt path, handling possible errors
func SetTxtPath() (bool, string) {
	var path string
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Please set path of .txt file (in such format: 'C:/Users/Admin/Desktop/file.txt'): ")
	scanner.Scan()
	path = scanner.Text()
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, "File doesn't exist."
		} else {
			return false, "Unknown error."
		}
	} else if filepath.Ext(strings.TrimSpace(path)) != ".txt" {
		return false, "Unsupported extension error."
	}
	return true, path
}

// Main menu
func main() {
	var firstOption int
	var path string
	flag.Parse()
	for {
		fmt.Println("Please select an option.\n" +
			"1. Set .txt path.\n" +
			"2. Count the number of lucky tickets.\n" +
			"3. Exit program.")
		for {
			_, err := fmt.Fscan(stdin, &firstOption)
			if err == nil {
				break
			}
			stdin.ReadString('\n')
			fmt.Println("Sorry, invalid input. Please enter correct number of option: ")
		}
		fmt.Println()
		switch firstOption {
		case 1:
			path = ""
			if isValid, str := SetTxtPath(); isValid {
				fmt.Println("Path has been set.")
				path = str
			} else {
				fmt.Println(str)
			}
			fmt.Println()
		case 2:
			moscowAmountOfLuckyTickets, petersburgAmountOfLuckyTickets, totalAmountOfLuckyTickets := FindLuckyTickets(path) // Output found lucky tickets up to the last line
			fmt.Printf("Number of Moscow lucky tickets: %d\n", moscowAmountOfLuckyTickets)                                  // of .txt file or up to the first error
			fmt.Printf("Number of Petersburg lucky tickets: %d\n", petersburgAmountOfLuckyTickets)
			fmt.Printf("Total number of lucky tickets: %d\n", totalAmountOfLuckyTickets)
			fmt.Println()
		case 3:
			os.Exit(3)
		default:
			fmt.Println("Unknown option.")
		}
	}
}
