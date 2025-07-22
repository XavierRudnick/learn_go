package main

import "fmt"

// add is a top-level function
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

//


var f func(x int, y int) int = func(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(f(42, 13)) // Call f, not add
}

//

func add(x int, y int) int {
	return x + y
}

// f is a package-level variable that holds the add function
var f func(x int, y int) int = add

func main() {
	fmt.Println(f(42, 13)) // Call f, which now points to add
}