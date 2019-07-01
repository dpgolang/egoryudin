package main

import (
	"Sixth/validation"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Поиск Московских счастливых билетов
func searchAllMoscow(ticketNumber int) bool {
	var rightSum int
	var leftSum int
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

// Поиск Питерских счастливых билетов
func searchAllPetersburg(ticketNumber int) bool {
	var evenSum int
	var oddSum int
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

// Поиск счастливых билетов
func FindLuckyTickets(path string) (int, int, int) {
	file, err := os.Open(path)
	rowCounter := 1                     // счётчик строк в файле. По большому счёту, создан для того чтобы информативнее показывать где был некорректный ввод
	totalAmountOfLuckyTickets := 0      // кол-во найденных счастливых билетов
	moscowAmountOfLuckyTickets := 0     // кол-во найденных счастилвых Московских билетов
	petersburgAmountOfLuckyTickets := 0 // кол-во найденных счастилвых Питерскихы билетов
	if err != nil {
		fmt.Println("Unable to read file.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // Построчное считываение файла.
		cityAndNumber := make([]string, 2)
		row := scanner.Text()
		cityAndNumber = strings.Split(row, "-")
		if len(cityAndNumber) != 2 {
			fmt.Printf("Error occurred while trying to read %d row\n"+
				"Fill .txt file in a such way: 'Moscow - 123123'\n\n", rowCounter)
			return moscowAmountOfLuckyTickets, petersburgAmountOfLuckyTickets, totalAmountOfLuckyTickets
		}
		for i := 0; i < len(cityAndNumber); i++ {
			cityAndNumber[i] = stripSpaces(cityAndNumber[i])
		}

		isCityCorrect, moscowOrPetersburg := validation.ValidateCity(cityAndNumber[0])
		isNumberChecked, ticketNumber := validation.ValidateInt(cityAndNumber[1])

		if isCityCorrect && isNumberChecked {

			if moscowOrPetersburg && searchAllMoscow(ticketNumber) {
				moscowAmountOfLuckyTickets++
				totalAmountOfLuckyTickets++
			} else if !moscowOrPetersburg && searchAllPetersburg(ticketNumber) {
				petersburgAmountOfLuckyTickets++
				totalAmountOfLuckyTickets++

			}
		}
		rowCounter++
	}
	return moscowAmountOfLuckyTickets, petersburgAmountOfLuckyTickets, totalAmountOfLuckyTickets
}
