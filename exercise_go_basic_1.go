package main

import (
	"fmt"
	"math"
)

func calculateCircumference(r int) float64 {
	return 2 * float64(r) * math.Pi
}

func roundCircumference(c float64) int {
	intPart := int(c)
	fracPart := c - float64(intPart)

	if fracPart >= 0.5 {
		return intPart + 1
	}
	return intPart
}

func main() {
	c := calculateCircumference(4)
	fmt.Println(roundCircumference(c))
}
