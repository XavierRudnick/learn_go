package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Println(i, f, b, s,"\n")
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", 3.14, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
"""a
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:

i := 42
f := float64(i)
u := uint(f)
"""