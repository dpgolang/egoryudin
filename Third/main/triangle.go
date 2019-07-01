package main

import "math"

type Triangle struct {
	name       string
	a, b, c, s float64
}

func (t *Triangle) HeronsFormula() {
	p := 0.5 * (t.a + t.b + t.c)
	t.s = math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}
