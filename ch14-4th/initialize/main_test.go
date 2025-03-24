package main

import "testing"

// goos: darwin
// goarch: arm64
// pkg: cty.sh/initialize
// cpu: Apple M2 Pro
// BenchmarkNew-10       11606             75733 ns/op
// BenchmarkAppend-10         77721            133324 ns/op
// PASS
// ok      cty.sh/initialize       26.779s

var t []int

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = InitSliceNew(i)
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = InitSliceAppend(i)
	}
}