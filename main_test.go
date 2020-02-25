package main

import (
	"math"
	"testing"
)

func Test_newBracket(t *testing.T) {

	b1 := newBracket(10000, 0, 0)
	b2 := newBracket(30000, 10000, .1)
	b3 := newBracket(100000, 30000, .25)

	if b1.totalTax != 0 || b2.totalTax != 2000 || b3.totalTax != 17500 {
		t.Error("Expected b1.totalTax = 0, got: ", b1.totalTax,
			"\nExpected b2.totalTax = 2000, got: ", b2.totalTax,
			"\n Expected b2.totalTax = 17500, got: ", b3.totalTax)
	}
}

func Test_calculateTax(t *testing.T) {

	top := math.Inf(0)

	b1 := newBracket(10000, 0, 0)
	b2 := newBracket(30000, 10000, .1)
	b3 := newBracket(100000, 30000, .25)
	b4 := newBracket(top, 100000, .4)

	brackets := make([]bracket, 4)
	brackets[0] = *b1
	brackets[1] = *b2
	brackets[2] = *b3
	brackets[3] = *b4

	t0 := calculateTax(42000, brackets)
	if t0 != 5000 {
		t.Error("Expected t0 = 5000, got: ", t0)
	}

	t1 := calculateTax(20000, brackets)
	if t1 != 1000 {
		t.Error("Expectd t1 = 1000, got: ", t1)
	}

	t2 := calculateTax(105000, brackets)
	if t2 != 21500 {
		t.Error("Expeted t2 = 21500, got ", t2)
	}
}
