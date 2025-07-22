package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	res := make(map[string]int)

	for _, v := range fields {
		_, ok := res[v]
		
		if ok == true {
			res[v] += 1
		} else {
			res[v] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
