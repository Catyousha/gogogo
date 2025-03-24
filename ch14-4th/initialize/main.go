package main

import (
	"fmt"
)

func InitSliceNew(n int) []int {
	s := make([]int, n)
	for i := range n {
		s[i] = i
	}
	return s
}

func InitSliceAppend(n int) []int {
	s := make([]int, 0)
	for i := range n {
		s = append(s, i)
	}
	return s
}

func main() {
	fmt.Println(InitSliceNew(10))
	fmt.Println(InitSliceAppend(10))
}
