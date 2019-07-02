package main

import (
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath  = filepath.Dir(b) + "\\file.txt"
)

type testTicketNumber struct {
	ticketNumber uint
	isCorrect    bool
}

var testsMoscow = []testTicketNumber{
	{123123, true},
	{321321, true},
	{113456, false},
	{234567, false},
}

var testsPetersburg = []testTicketNumber{
	{111122, true},
	{131133, false},
	{123123, false},
	{172222, true},
}

type testLuckyTickets struct {
	path string
	totalTickets, moscowTickets, petersburgTickets uint
}

var testsLuckyTickets = []testLuckyTickets{
	{basepath, 2, 1, 1},
}

func TestSearchAllMoscow(t *testing.T) {
	for _, pair := range testsMoscow {
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
	for _, pair := range testsPetersburg {
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

func TestFindLuckyTickets(t *testing.T) {
	for _, pair := range testsLuckyTickets {
		t1, t2, t3 := FindLuckyTickets(pair.path)
		if t1 != pair.moscowTickets || t2 != pair.petersburgTickets || t3 != pair.totalTickets{
			t.Error(
				"For", pair.path,
				"expected", pair.moscowTickets, pair.petersburgTickets, pair.totalTickets,
				"got", t1, t2, t3,
			)
		}
	}
}

