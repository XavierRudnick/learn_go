package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for math.Abs(x-z*z) > .05{ 
		z -= (z*z - x) / (2*z)
		println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(16))
}
