package validation

import (
	"fmt"
	"strconv"
)

const moscowAlgorithm string = "Moscow"
const petersburgAlgorithm string = "Petersburg"

func ValidateNumber(s string) (uint, error) {
	if len(s) == 6 {
		number, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("error occured: %s", err)
		}
		return uint(number), nil
	}
	return 0, fmt.Errorf("error occured: ticket number '%s' must be a length of 6", s)
}

// Checking whether first value of the row (city) was input correctly. Defining an algorithm
func ValidateCity(s string) (string, error) {
	if s == moscowAlgorithm {
		return moscowAlgorithm, nil
	} else if s == petersburgAlgorithm {
		return petersburgAlgorithm, nil
	}
	return  "", fmt.Errorf("error occured: cannot define an algorithm '%s'", s)
}

// Each row of the file must consist of 2 values: name of the city (algorithm) and ticket number
func ValidateRow(s []string) error {
	if len(s) != 2 {
		return fmt.Errorf("error occured: two parameters needed")
	}
	return nil
}
