package main

import (
	"fmt"
	"os"
	"strconv"
)

type ar2x2 [2][2]int

// Traditional Add() function
func Add(a, b ar2x2) ar2x2 {
	c := ar2x2{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

func (a *ar2x2) Add(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = a[i][j] + b[i][j]
		}
	}
}

func (a *ar2x2) Subtract(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = a[i][j] - b[i][j]
		}
	}
}

func (a *ar2x2) Multiply(b ar2x2) {
	a[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0]
	a[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0]
	a[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1]
	a[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1]
}

func main() {
	// go run . 1 2 3 4 5 6 7 8
	if len(os.Args) != 9 {
		fmt.Println("Need 8 integers")
		return
	}

	k := [8]int{}
	for index, i := range os.Args[1:] {
		v, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		k[index] = v
	}
	// 1 2 3 4
	a := ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	// 5 6 7 8
	b := ar2x2{{k[4], k[5]}, {k[6], k[7]}}

	// Traditional a+b: [[6 8] [10 12]]
	fmt.Println("Traditional a+b:", Add(a, b))

	// a+b: [[6 8] [10 12]]
	a.Add(b) // a is modified
	fmt.Println("a+b:", a)

	// a-a: [[0 0] [0 0]]
	a.Subtract(a)
	fmt.Println("a-a:", a)

	// a*b: [[19 130] [43 290]]
	a = ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	a.Multiply(b)
	fmt.Println("a*b:", a)

	// b*a: [[23 70] [31 94]]
	a = ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	b.Multiply(a)
	fmt.Println("b*a:", b)
}
