package validation

import (
	"errors"
	"strings"
	"testing"
)

type testCity struct {
	city string
	err error
}

var testsCity = []testCity{
	{"Moscow", nil},
	{"Petersburg", nil},
	{"Dnipro", errors.New("cannot define an algorithm")},
	{"Lviv", errors.New("cannot define an algorithm")},
	{"1", errors.New("cannot define an algorithm")},
}

type testNumber struct {
	number string
	err error
}

var testsNumbers = []testNumber{
	{"123123", nil},
	{"12W412", errors.New("is not an integer number")},
	{"-12344", errors.New("is not an integer number")},

}

type testRow struct {
	elems []string
	err error
}

var testsRow = []testRow {
	{[]string{"Moscow", "123123"}, nil},
	{[]string{"Petersburg"}, errors.New("two parameters needed")},
	{[]string{"Moscow", "543123"}, nil},
}

func TestValidateCity(t *testing.T) {
	for _, pair := range testsCity {
		_, err := ValidateCity(pair.city)
		if (err != nil && !strings.Contains(err.Error(), pair.err.Error())) || (err == nil && err != pair.err) {
			t.Error(
				"For", pair.city,
				"expected", pair.err,
				"got", err,
			)
		}
	}
}

func TestValidateNumber(t *testing.T) {
	for _, pair := range testsNumbers {
		_, err := ValidateNumber(pair.number)
		if (err != nil && !strings.Contains(err.Error(), pair.err.Error())) || (err == nil && err != pair.err) {
			t.Error(
				"For", pair.number,
				"expected", pair.err,
				"got", err,
			)
		}
	}
}

func TestValidateRow(t *testing.T) {
	for _, pair := range testsRow {
		err := ValidateRow(pair.elems)
		if (err != nil && !strings.Contains(err.Error(), pair.err.Error())) || (err == nil && err != pair.err) {
			t.Error(
				"For", pair.elems,
				"expected", pair.err,
				"got", err,
			)
		}
	}
}
