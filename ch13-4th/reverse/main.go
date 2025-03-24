package main

// This version has bugs

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func R1(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("Invalid UTF-8")
	}
	a := []byte(s)
	for i, j := 0, len(s)-1; i < j; i++ {
		a[i], a[j] = a[j], a[i]
		j--
	}
	return string(a), nil
}

func R2(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("Invalid UTF-8")
	}
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b), nil
}

func main() {
	str := "1234567890"
	R1_r, _ := R1(str)
	fmt.Println(string(R1_r))
	
	fmt.Println(R2(str))
}
