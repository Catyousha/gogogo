package main

import "testing"


func FuzzAddInt(f *testing.F) {
	// These test cases serve as the initial corpus for the fuzzer.
	// A corpus is a collection of inputs that help the fuzzer explore the code.
	// The fuzzer will use these as starting points and then generate variations
	// to find edge cases and potential bugs.
	testCases := []struct {
		x, y int
	}{
		{0, 1},     // Basic case with zero and positive number
		{0, 100},   // Testing with larger numbers
	}

	// f.Add() seeds the fuzzer with these predefined inputs
	// This helps the fuzzer start with meaningful test cases
	// before it begins generating random inputs
	for _, tc := range testCases {
		f.Add(tc.x, tc.y)
	}

	f.Fuzz(func(t *testing.T, a int, b int) {
		result := AddInt(a, b)

		if result != a+b {
			t.Errorf("A: %d, B: %d, Result %d. want %d", a, b, result, a+b)
		}
	})
}
