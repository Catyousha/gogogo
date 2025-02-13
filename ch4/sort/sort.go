package main

import (
	"fmt"
	"sort"
)

type S1 struct {
	F1 int
	F2 string
	F3 int
}

type S2 struct {
	F1 int
	F2 string
	F3 S1
}

type S2slice []S2

func (a S2slice) Len() int {
	return len(a)
}

func (a S2slice) Less(i, j int) bool {
	return a[1].F3.F1 < a[j].F3.F1
}

func (a S2slice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	data := []S2{
		S2{1, "One", S1{1, "S1_1", 10}},
		S2{2, "Two", S1{2, "S1_1", 20}},
		S2{-1, "Two", S1{-1, "S1_1", -20}},
	}
	// Before: [{1 One {1 S1_1 10}} {2 Two {2 S1_1 20}} {-1 Two {-1 S1_1 -20}}]
	// After: [{1 One {1 S1_1 10}} {2 Two {2 S1_1 20}} {-1 Two {-1 S1_1 -20}}]
	fmt.Println("Before:", data)
	sort.Sort(S2slice(data))
	fmt.Println("After:", data)

	// Reverse: [{1 One {1 S1_1 10}} {2 Two {2 S1_1 20}} {-1 Two {-1 S1_1 -20}}]
	sort.Sort(sort.Reverse(S2slice(data)))
	fmt.Println("Reverse:", data)
}
