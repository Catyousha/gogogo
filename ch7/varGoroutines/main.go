package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	count := 10
	args := os.Args
	if len(args) == 2 {
		t, err := strconv.Atoi(args[1])
		if err == nil {
			count = t
		}
	}

	fmt.Println(count, "Goroutines:")
	var wg sync.WaitGroup
	// sync.WaitGroup{noCopy:sync.noCopy{}, state:atomic.Uint64{_:atomic.noCopy{}, _:atomic.align64{}, v:0x0}, sema:0x0}
	fmt.Printf("%#v\n", wg)

	// 9
	// 6
	// 7
	// 8
	// 0
	// 3
	// 4
	// 1
	// 5
	// 2
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Println(x)
		}(i)
	}

	// sync.WaitGroup{noCopy:sync.noCopy{}, state:atomic.Uint64{_:atomic.noCopy{}, _:atomic.align64{}, v:0xa00000000}, sema:0x0}
	fmt.Printf("%#v\n", wg)
	wg.Wait()
}