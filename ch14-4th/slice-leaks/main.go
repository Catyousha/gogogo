package main

import (
	"fmt"
	"time"
)

func createSlice() []int {
	return make([]int, 1000000)
}

func getValueLeak(s []int) []int {
	val := s[:3]
	return val
}


func sliceLeaks() {
	for range 15 {
		message := createSlice()
		val := getValueLeak(message)
		fmt.Print(len(val), " ")
		time.Sleep(10 * time.Millisecond)
	}
}


func getValueNoLeak(s []int) []int {
	returnVal := make([]int, 3)
	copy(returnVal, s)
	return returnVal
}

func sliceNoLeaks() {
	for range 15 {
		message := createSlice()
		val := getValueNoLeak(message)
		fmt.Print(len(val), " ")
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// go run -gcflags '-m -l' .
	// ./main.go:9:13: make([]int, 1000000) escapes to heap
	// ./main.go:12:19: leaking param: s to result ~r0 level=0
	// ./main.go:22:12: ... argument does not escape
	// ./main.go:22:16: len(val) escapes to heap
	// ./main.go:22:23: " " escapes to heap
	// 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3
	sliceLeaks()


	// ./main.go:9:13: make([]int, 1000000) escapes to heap
	// ./main.go:28:21: s does not escape
	// ./main.go:29:19: make([]int, 3) escapes to heap
	// ./main.go:38:12: ... argument does not escape
	// ./main.go:38:16: len(val) escapes to heap
	// ./main.go:38:23: " " escapes to heap
	// 3 3 3 3 3 3 3 3 3 3 3 3 3 3 3
	sliceNoLeaks()
}