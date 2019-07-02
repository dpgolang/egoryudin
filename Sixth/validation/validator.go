package validation

import (
	"fmt"
	"strconv"
)

const moscowAlgorithm string = "Moscow"
const petersburgAlgorithm string = "Petersburg"

func ValidateNumber(s string) (error, uint) {
	if len(s) == 6 {
		number, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return fmt.Errorf("error occured: %s", err), 0
		}
		return nil, uint(number)
	}
	return  fmt.Errorf("error occured: ticket number '%s' must be a length of 6", s), 0
}

// Checking whether first value of the row (city) was input correctly. Defining an algorithm
func ValidateCity(s string) (error, string) {
	if s == moscowAlgorithm {
		return nil, moscowAlgorithm
	} else if s == petersburgAlgorithm {
		return nil, petersburgAlgorithm
	}
	return fmt.Errorf("error occured: cannot define an algorithm '%s'", s), ""
}

// Each row of the file must consist of 2 values: name of the city (algorithm) and ticket number
func ValidateRow(s []string) error {
	if len(s) != 2 {
		return fmt.Errorf("error occured: two parameters needed")
	}
	return nil
}
