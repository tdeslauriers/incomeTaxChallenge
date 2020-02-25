package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	inc := os.Args[1]
	i, err := strconv.ParseFloat(inc, 64)
	if err != nil {
		panic("not a number.")
	}

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

	fmt.Printf("Total Tax owed for income of %.2f is %.2f.\n", i, calculateTax(i, brackets))

}

type bracket struct {
	income   float64
	top      float64
	bottom   float64
	taxRate  float64
	totalTax float64
}

func newBracket(top float64, bottom float64, taxRate float64) *bracket {

	b := bracket{top: top, bottom: bottom, taxRate: taxRate}
	if top < math.Inf(0) {

		b.totalTax = (top - bottom) * taxRate
	} else {

		b.totalTax = (b.income - bottom) * taxRate
	}

	return &b
}

func calculateTax(income float64, braks []bracket) float64 {

	var total float64

	for i := 0; i < len(braks); i++ {

		if income > braks[i].top {

			total += braks[i].totalTax

		} else if income < braks[i].top {

			final := (income - braks[i].bottom) * braks[i].taxRate
			total += final

			return total
		}
	}

	return total
}
