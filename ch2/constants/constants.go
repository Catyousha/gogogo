package main

import "fmt"

type Digit int
type Power2 int

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const (
		Zero = iota * 2
		One
		Two
	)

	fmt.Println(One)
	fmt.Println(Two)

	const (
		p2_0 Power2 = 1 << iota
		p2_1
		p2_2
	)

	fmt.Println("2^1:",p2_1);
	fmt.Println("2^2:",p2_2);
}
