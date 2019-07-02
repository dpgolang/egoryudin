package main

import (
	"Sixth/validation"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Moscow algorithm
func searchAllMoscow(ticketNumber uint) bool {
	var rightSum uint
	var leftSum uint
	for i := 0; i < 6; i++ {
		digit := ticketNumber % 10
		ticketNumber /= 10
		if i < 3 {
			rightSum += digit
		} else {
			leftSum += digit
		}
	}
	return rightSum == leftSum
}

// Petersburg algorithm
func searchAllPetersburg(ticketNumber uint) bool {
	var evenSum uint
	var oddSum uint
	for i := 0; i < 6; i++ {
		digit := ticketNumber % 10
		ticketNumber /= 10
		if digit&1 == 0 {
			evenSum += digit
		} else {
			oddSum += digit
		}
	}
	return evenSum == oddSum
}

// Search lucky tickets
func FindLuckyTickets(path string) (uint, uint, uint) {
	var (
		rowCounter                     = 1 // .txt row counter. Actually, this variable was created in order to more informatively show where an incorrect input happened
		totalAmountOfLuckyTickets      uint
		moscowAmountOfLuckyTickets     uint
		petersburgAmountOfLuckyTickets uint
	)
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Unable to read file.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // line by line .txt file reading
		cityAndNumber := make([]string, 2)
		row := scanner.Text()
		cityAndNumber = strings.Split(row, "-")

		if rowError := validation.ValidateRow(cityAndNumber); rowError != nil {
			fmt.Printf("Bad input at %d row, %s\n", rowCounter, rowError)
			return moscowAmountOfLuckyTickets, petersburgAmountOfLuckyTickets, totalAmountOfLuckyTickets
		}

		// Removing all spaces from strings
		for i := 0; i < len(cityAndNumber); i++ {
			cityAndNumber[i] = stripSpaces(cityAndNumber[i])
		}

		moscowOrPetersburg, cityError := validation.ValidateCity(cityAndNumber[0])
		if cityError != nil {
			fmt.Printf("Bad input at %d row, %s\n", rowCounter, cityError)
			return moscowAmountOfLuckyTickets, petersburgAmountOfLuckyTickets, totalAmountOfLuckyTickets
		}

		ticketNumber, numberError := validation.ValidateNumber(cityAndNumber[1])
		if numberError != nil {
			fmt.Printf("Bad input at %d row, %s\n", rowCounter, numberError)
			return moscowAmountOfLuckyTickets, petersburgAmountOfLuckyTickets, totalAmountOfLuckyTickets
		}

		if moscowOrPetersburg == "Moscow" && searchAllMoscow(ticketNumber) {
			moscowAmountOfLuckyTickets++
			totalAmountOfLuckyTickets++
		} else if moscowOrPetersburg == "Petersburg" && searchAllPetersburg(ticketNumber) {
			petersburgAmountOfLuckyTickets++
			totalAmountOfLuckyTickets++
		}
		rowCounter++
	}
	return moscowAmountOfLuckyTickets, petersburgAmountOfLuckyTickets, totalAmountOfLuckyTickets
}
