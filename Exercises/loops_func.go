package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func Sqrt2(x, eps float64) float64 {
	z1, z2 := 1.0, 1.0
	diff := eps
	for diff >= eps {
		z2 = z1
		z1 -= (z1*z1 - x) / (2 * z1)
		diff = math.Abs(z1 - z2)
	}
	return z1
}

func main() {
	var v float64 = 6
	fmt.Println(Sqrt(v))
	fmt.Println(Sqrt2(v, 0.000000005))
}