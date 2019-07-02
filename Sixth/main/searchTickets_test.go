package main

import "testing"

type testpair struct {
	ticketNumber uint
	isCorrect bool
}

var testsMoscow = []testpair{
	{123123, true},
	{321321, true},
	{113456, false},
	{234567, false},
}

var testsPetersburg = []testpair{
	{111122, true},
	{131133, false},
	{123123, false},
	{172222, true},
}

func TestSearchAllMoscow(t *testing.T) {
	for _ , pair := range testsMoscow {
		v := searchAllMoscow(pair.ticketNumber)
		if v != pair.isCorrect {
			t.Error(
				"For", pair.ticketNumber,
				"expected", pair.isCorrect,
				"got", v,
			)
		}
	}
}

func TestSearchAllPetersburg(t *testing.T) {
	for _ , pair := range testsPetersburg {
		v := searchAllPetersburg(pair.ticketNumber)
		if v != pair.isCorrect {
			t.Error(
				"For", pair.ticketNumber,
				"expected", pair.isCorrect,
				"got", v,
			)
		}
	}
}