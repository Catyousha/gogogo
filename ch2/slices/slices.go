package main

import (
	"fmt"
)

// exercise 1
func concatenateArraysToSlice(arr1, arr2 [3]int) []int {
    slice1 := arr1[:]
    slice2 := arr2[:]

	return append(slice1, slice2...)
}

// exercise 2
func concatenateArrays(arr1 [3]int, arr2 [4]int) [7]int {
    var result [7]int
    copy(result[:], arr1[:]) 

    for i := 0; i < len(arr2); i++ {
        result[len(arr1)+i] = arr2[i]
    }

    return result
}

// exercise 3
func concatenateSlicesToArray(slice1, slice2 []int) []int {
	result := make([]int, len(slice1)+len(slice2))
    copy(result[:], slice1[:]) 

    for i := 0; i < len(slice2); i++ {
        result[len(slice1)+i] = slice2[i]
    }
    return result
}


func main() {
	// cap = slice capacity / preserved memory
	// [0 1 2 3 4] | len: 5 | cap: 5
	aSlice := []int{0, 1, 2, 3, 4}
	fmt.Println(aSlice, "| len:", len(aSlice), "| cap:", cap(aSlice))

	// it doubles!
	// [0 1 2 3 4 5] | len: 6 | cap: 10
	aSlice = append(aSlice, 5)
	fmt.Println(aSlice, "| len:", len(aSlice), "| cap:", cap(aSlice))

	// optimize memory
	// [0 1 2 3 4 5] | len: 6 | cap: 6
	l := len(aSlice)
	aSlice = aSlice[0:l:l]
	fmt.Println(aSlice, "| len:", len(aSlice), "| cap:", cap(aSlice))

	// [0 1 2 3 4 5]
	arr1 := [3]int{0, 1, 2}
	arr2 := [3]int{3, 4, 5}
	fmt.Println(concatenateArraysToSlice(arr1, arr2));

	// [1 2 3 4 5 6 7]
	arr3 := [3]int{1,2,3}
	arr4 := [4]int{4,5,6,7}
	fmt.Println(concatenateArrays(arr3, arr4))

	// [1 2 3 4 5 6]
	slice1 := []int{1,2,3}
	slice2 := []int{4,5,6}
	fmt.Println(concatenateSlicesToArray(slice1, slice2))
}
