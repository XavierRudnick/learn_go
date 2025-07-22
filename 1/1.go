package main //all programs made of packages

import (
	"fmt"
	"math/rand" //last name in path is package referr
	"math"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi) // exported names like Pi need to be capitalized
}
