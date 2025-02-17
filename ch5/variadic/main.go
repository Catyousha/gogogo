package main

import (
	"fmt"
	"os"
)

func addFloats(message string, s ...float64) float64 {
	fmt.Println(message)
	sum := float64(0)
	for _, a := range s {
		sum = sum + a
	}
	s[0] = -1000
	return sum;
}

func everything(input ...interface{}) {
	fmt.Println(input...)
}

func main() {
	// Adding numbers...
	// Sum: 24.36
	sum := addFloats("Adding numbers...", 1.1, 2.12, 3.14, 4, 5, -1, 10)
	fmt.Println("Sum:", sum)

	// Adding numbers:
	// Sum: 6.36
	s := []float64{1.1, 2.12, 3.14}
	sum = addFloats("Adding numbers:", s...)
	fmt.Println("Sum:", sum)

	// [-1000 2.12 3.14]
	everything(s)

	// go run . 1 21 35
	// [1 21 35]
	empty := make([]interface{}, len(os.Args[1:]))
	for i, v := range os.Args[1:] {
		empty[i] = v
	}
	everything(empty)
}