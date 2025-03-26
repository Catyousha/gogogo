package main

import (
	"fmt"
	"runtime"
)

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}

func main()  {
	// 129 KB
	n := 2_000_000
	m := make(map[int][128]byte)
	printAlloc()

	// alloc n * 128 bytes
	// 745840 KB
	for i := range n {
		m[i] = [128]byte{}
	}
	printAlloc()

	// 745841 KB
	for i := range n {
		delete(m, i)
	}
	printAlloc()

	// 588703 KB
	runtime.GC()
	printAlloc()

	// 160 KB
	runtime.KeepAlive(m)
	m = nil
	runtime.GC()
	printAlloc()
}