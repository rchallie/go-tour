package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 3.0
	var z_prev float64
	for i := 0; z != z_prev && i < 10; i++ {
		z_prev = z
		z -= (z*z - x) / (2*z)
		fmt.Println("z :=", z)
	}
	return z
}

func main() {
		fmt.Println("Mine:", Sqrt(2))
		fmt.Println("Math:", math.Sqrt(2))
}
