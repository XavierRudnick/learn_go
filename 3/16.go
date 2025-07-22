package main                                                                                                   

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//	var primes = [6]int{2, 3, 5, 7, 11, 13}
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]//include exclude
	fmt.Println(primes)
	fmt.Println(s)

}
