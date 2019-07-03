package validation

import (
	"testing"
)

type boolSlice []bool

func (b boolSlice) contains() bool {
	for _, a := range b {
		if !a {
			return true
		}
	}
	return false
}

type testExists struct {
	sides     []float64
	isCorrect bool
}

var testsExists = []testExists{
	{[]float64{5, 6, 10}, true},
	{[]float64{5, 6, 11}, false},
	{[]float64{1, 2, 3}, false},
}

type testFloat struct {
	number    string
	isCorrect bool
}

var testsFloat = []testFloat{
	{"1", true},
	{"-1", false},
	{"14.23", true},
	{"w", false},
}

type testAnswer struct {
	answer      string
	yesOrNo     bool
	isIncorrect bool
}

var testsAnswer = []testAnswer{
	{"yes", true, false},
	{"n", false, false},
	{"awd", false, true},
}

type testFinalTriangle struct {
	splitSlice []string
	isCorrect  []bool
}

var testsFinalTriangle = []testFinalTriangle{
	{[]string{"first", "5", "6", "10"}, []bool{true, true, true, true}},
	{[]string{"second", "5", "6", "11"}, []bool{true, false, false, false}},
	{[]string{"third", "-1", "6", "10"}, []bool{true, false, true, true}},
	{[]string{"fourth", "5", "w", "10"}, []bool{true, true, false, true}},
}

func TestTriangleExistsValidator(t *testing.T) {
	for _, pair := range testsExists {
		res := TriangleExistsValidator(pair.sides)
		if pair.isCorrect != res {
			t.Error(
				"For", pair.sides,
				"expected", pair.isCorrect,
				"got", res,
			)
		}
	}
}

func TestValidateFloat(t *testing.T) {
	for _, pair := range testsFloat {
		_, res := ValidateFloat(pair.number)
		if pair.isCorrect != res {
			t.Error(
				"For", pair.number,
				"expected", pair.isCorrect,
				"got", res,
			)
		}
	}
}

func TestValidateAnswer(t *testing.T) {
	for _, pair := range testsAnswer {
		res := ValidateAnswer(pair.answer)
		if pair.yesOrNo != res || pair.isIncorrect != IsAnswerIncorrect() {
			t.Error(
				"For", pair.answer,
				"expected", pair.yesOrNo,
				"got", res,
			)
		}
	}
}

func TestFinalTriangleValidator(t *testing.T) {
	for _, pair := range testsFinalTriangle {
		_, res := FinalTriangleValidator(pair.splitSlice)
		expect := boolSlice(pair.isCorrect).contains()
		act := boolSlice(res).contains()
		if act != expect {
			t.Error(
				"For", pair.splitSlice,
				"expected", expect,
				"got", act,
			)
		}
	}
}
