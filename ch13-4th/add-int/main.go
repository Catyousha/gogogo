package main

import (
	"fmt"
)

func AddInt(x, y int) int {
	// Exercise 3: Fix the AddInt() function from fuzz/code.go.
	if(x < 0) {
		for range x *-1 {
			y = y-1
		}
		return y
	}
	for range x {
		y = y + 1
	}
	return y
}
func main() {
	fmt.Println(AddInt(5, 4))
}
