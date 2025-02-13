package main

import (
	"fmt"
	"math"
)

type Shape2D interface {
	Perimeter() float64
}

type circle struct {
	R float64
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func main()  {
	a := circle{R: 1.5}
	// R  2 -> Perimeter   9
	fmt.Printf("R %2.f -> Perimeter %3.f \n", a.R, a.Perimeter())

	_, ok := interface{}(a).(Shape2D)
	if ok {
		fmt.Println("a is a Shape2D!")
	}
}