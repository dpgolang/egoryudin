package main

import "strconv"

var (
	digits = []string{
		1: "один",
		"два",
		"три",
		"четыре",
		"пять",
		"шесть",
		"семь",
		"восемь",
		"девять",
	}

	digitsExcep = []string{
		"одна",
		"две",
	}

	teens = []string{
		1: "одинадцать",
		"двенадцать",
		"тринадцать",
		"четырнадцать",
		"пятнадцать",
		"шестнадцать",
		"семнадцать",
		"восемнадцать",
		"девятнадцать",
	}

	tens = []string{
		1: "десять",
		"двадцать",
		"тридцать",
		"сорок",
		"пятьдесят",
		"шестьдесят",
		"семьдесят",
		"восемьдесят",
		"девяносто",
	}

	hundreds = []string{
		1: "сто",
		"двести",
		"триста",
		"четыреста",
		"пятьсот",
		"шестьсот",
		"семьсот",
		"восемьсот",
		"девятьсот",
	}

	thousands = []string{
		"тысяча",
		"тысячи",
		"тысяч",
	}

	millions = []string{
		"миллион",
		"миллиона",
		"миллионов",
	}

	billions = []string{
		"миллиард",
		"миллиарда",
		"миллиардов",
	}
)

func getDigitNumString(number int64) string {
	return digits[number]
}

func getTenOrTeenNumString(number int64) string {
	if number < 20 {
		if number == 10 {
			return tens[1]
		}
		return teens[number%10]
	} else if number%10 == 0 {
		return tens[number/10]
	}
	var (
		output         string
		splittedDigits [2]int64
	)
	// filling splittedDigits array backwards
	for i := 1; i > -1; i-- {
		splittedDigits[i] = number % 10
		number /= 10
	}
	output = tens[splittedDigits[0]] + " " + digits[splittedDigits[1]]
	return output
}

func getHundredNumString(number int64) string {
	var output string
	remainder := number % 100
	hundredStr := hundreds[number/100]
	if remainder == 0 {
		output = hundredStr
	} else if remainder < 10 {
		output = hundredStr + " " + getDigitNumString(remainder)
	} else {
		output = hundredStr + " " + getTenOrTeenNumString(remainder)
	}
	return output
}

func getThousandNumString(number int64) string {
	var output string
	remainder := number % 1000
	thousandStr := getThousandString(number / 1000)
	if remainder == 0 {
		output = thousandStr
	} else if remainder > 0 && remainder < 10 {
		output = thousandStr + " " + getDigitNumString(remainder)
	} else if remainder >= 10 && remainder < 100 {
		output = thousandStr + " " + getTenOrTeenNumString(remainder)
	} else {
		output = thousandStr + " " + getHundredNumString(remainder)
	}
	return output
}

func getThousandString(thousandNumber int64) string {
	var output string
	switch len(strconv.FormatInt(thousandNumber, 10)) {
	case 1:
		output = getOneDigitThousandString(thousandNumber)
	case 2:
		output = getTwoDigitsThousandString(thousandNumber)
	case 3:
		output = getThreeDigitsThousandString(thousandNumber)
	}
	return output
}

func getOneDigitThousandString(thousandNum int64) string {
	var output string
	if thousandNum == 1 {
		output = digitsExcep[0] + " " + thousands[0]
	} else if thousandNum == 2 {
		output = digitsExcep[1] + " " + thousands[1]
	} else if thousandNum == 3 || thousandNum == 4 {
		output = getDigitNumString(thousandNum) + " " + thousands[1]
	} else {
		output = getDigitNumString(thousandNum) + " " + thousands[2]
	}
	return output
}

func getTwoDigitsThousandString(thousandNum int64) string {
	var output string
	remainder := thousandNum % 10
	if thousandNum > 10 && thousandNum < 20 {
		output = getTenOrTeenNumString(thousandNum) + " " + thousands[2]
	} else if remainder == 1 {
		output = getTenOrTeenNumString(thousandNum-remainder) + " " + getOneThousandString()
	} else if remainder == 2 {
		output = getTenOrTeenNumString(thousandNum-remainder) + " " + getTwoThousandString()
	} else if remainder == 3 || remainder == 4 {
		output = getTenOrTeenNumString(thousandNum) + " " + thousands[1]
	} else {
		output = getTenOrTeenNumString(thousandNum) + " " + thousands[2]
	}
	return output
}

func getThreeDigitsThousandString(thousandNum int64) string {
	var output string
	remainderOnTen := thousandNum % 10
	remainderOnHundred := thousandNum % 100
	if remainderOnHundred > 10 && remainderOnHundred < 20 {
		output = getHundredNumString(thousandNum-remainderOnHundred) +
			" " + getTenOrTeenNumString(remainderOnHundred) + " " + thousands[2]
	} else if remainderOnTen == 1 {
		output = getHundredNumString(thousandNum-remainderOnTen) + " " + getOneThousandString()
	} else if remainderOnTen == 2 {
		output = getHundredNumString(thousandNum-remainderOnTen) + " " + getTwoThousandString()
	} else if remainderOnTen == 3 || remainderOnTen == 4 {
		output = getHundredNumString(thousandNum) + " " + thousands[1]
	} else {
		output = getHundredNumString(thousandNum) + " " + thousands[2]
	}
	return output
}

func getOneThousandString() string {
	return digitsExcep[0] + " " + thousands[0]
}

func getTwoThousandString() string {
	return digitsExcep[1] + " " + thousands[1]
}

func getMillionNumString(number int64) string {
	var output string
	remainder := number % 1000000
	millionStr := getMillionString(number / 1000000)
	if remainder == 0 {
		output = millionStr
	} else if remainder < 1000 {
		output = millionStr + " " + getStringEnding(int64(len(strconv.FormatInt(remainder, 10))), remainder)
	} else {
		thousandStr := getThousandNumString(remainder)
		output = millionStr + " " + thousandStr
	}
	return output
}

func getMillionString(millionNum int64) string {
	var output string
	switch len(strconv.FormatInt(millionNum, 10)) {
	case 1:
		output = getOneDigitMillionString(millionNum)
	case 2:
		output = getTwoDigitsMillionString(millionNum)
	case 3:
		output = getThreeDigitsMillionString(millionNum)
	}
	return output
}

func getStringEnding(length, num int64) string {
	var output string
	switch length {
	case 1:
		output = getDigitNumString(num)
	case 2:
		output = getTenOrTeenNumString(num)
	case 3:
		output = getHundredNumString(num)
	}
	return output
}

func getOneDigitMillionString(firstPartOfNum int64) string {
	var output string
	if firstPartOfNum == 1 {
		output = digits[1] + " " + millions[0]
	} else if firstPartOfNum == 2 {
		output = digits[2] + " " + millions[1]
	} else if firstPartOfNum == 3 || firstPartOfNum == 4 {
		output = getDigitNumString(firstPartOfNum) + " " + millions[1]
	} else {
		output = getDigitNumString(firstPartOfNum) + " " + millions[2]
	}
	return output
}

func getTwoDigitsMillionString(millionNum int64) string {
	var output string
	remainder := millionNum % 10
	if millionNum > 10 && millionNum < 20 {
		output = getTenOrTeenNumString(millionNum) + " " + millions[2]
	} else if remainder == 1 {
		output = getTenOrTeenNumString(millionNum) + " " + millions[0]
	} else if remainder >= 2 && remainder <= 4 {
		output = getTenOrTeenNumString(millionNum) + " " + millions[1]
	} else {
		output = getTenOrTeenNumString(millionNum) + " " + millions[2]
	}
	return output
}

func getThreeDigitsMillionString(millionNum int64) string {
	var output string
	remainderOnHundred := millionNum % 100
	remainderOnTen := millionNum % 10
	if remainderOnHundred > 10 && remainderOnHundred < 20 {
		output = getHundredNumString(millionNum) + " " + millions[2]
	} else if remainderOnTen == 1 {
		output = getHundredNumString(millionNum) + " " + millions[0]
	} else if remainderOnTen >= 2 && remainderOnTen <= 4 {
		output = getHundredNumString(millionNum) + " " + millions[1]
	} else {
		output = getHundredNumString(millionNum) + " " + millions[2]
	}
	return output
}

func getBillionNumString(number int64) string {
	remainderOnBillion := number % 1000000000
	billionStr := getBillionString(number / 1000000000)
	if remainderOnBillion == 0 {
		return billionStr
	} else if remainderOnBillion < 1000 {
		return billionStr + " " + getStringEnding(int64(len(strconv.FormatInt(remainderOnBillion, 10))), remainderOnBillion)
	} else if remainderOnBillion < 1000000 {
		return billionStr + " " + getThousandNumString(remainderOnBillion)
	} else {
		return billionStr + " " + getMillionNumString(remainderOnBillion)
	}
}

func getBillionString(billionNum int64) string {
	var output string
	switch int64(len(strconv.FormatInt(billionNum, 10))) {
	case 1:
		output = getOneDigitBillionString(billionNum)
		break
	case 2:
		output = getTwoDigitsBillionString(billionNum)
		break
	case 3:
		output = getThreeDigitsBillionString(billionNum)
		break
	}
	return output
}

func getOneDigitBillionString(billionNum int64) string {
	var output string
	if billionNum == 1 {
		output = digits[1] + " " + billions[0]
	} else if billionNum == 2 {
		output = digits[2] + " " + billions[1]
	} else if billionNum == 3 || billionNum == 4 {
		output = getDigitNumString(billionNum) + " " + billions[1]
	} else {
		output = getDigitNumString(billionNum) + " " + billions[2]
	}
	return output
}

func getTwoDigitsBillionString(billionNum int64) string {
	var output string
	remainderOnTen := billionNum % 10
	if billionNum > 10 && billionNum < 20 {
		output = getTenOrTeenNumString(billionNum) + " " + billions[2]
	} else if remainderOnTen == 1 {
		output = getTenOrTeenNumString(billionNum) + " " + billions[0]
	} else if remainderOnTen >= 2 && remainderOnTen <= 4 {
		output = getTenOrTeenNumString(billionNum) + " " + billions[1]
	} else {
		output = getTenOrTeenNumString(billionNum) + " " + billions[2]
	}
	return output
}

func getThreeDigitsBillionString(billionNum int64) string {
	var output string
	remainderOnHundred := billionNum % 100
	remainderOnTen := billionNum % 10
	if remainderOnHundred > 10 && remainderOnHundred < 20 {
		output = getHundredNumString(billionNum) + " " + billions[2]
	} else if remainderOnTen == 1 {
		output = getHundredNumString(billionNum) + " " + billions[0]
	} else if remainderOnTen >= 2 && remainderOnTen <= 4 {
		output = getHundredNumString(billionNum) + " " + billions[1]
	} else {
		output = getHundredNumString(billionNum) + " " + billions[2]
	}
	return output
}
