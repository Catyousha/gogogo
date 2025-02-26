package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex
var v1 int

func change() {
	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1 == 10 {
		v1 = 0
		fmt.Print("* \n")
	}
	m.Unlock()
}

func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func main()  {
	num := 25

	var wg sync.WaitGroup
	// -> 1-> 2-> 3-> 4-> 5-> 6-> 7-> 8-> 9* 
	// -> 0-> 1-> 2-> 3-> 4-> 5-> 6-> 7-> 8-> 9* 
	// -> 0-> 1-> 2-> 3-> 4-> 5
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int)  {
			defer wg.Done()
			change()
			fmt.Printf("-> %d", read())
		}(i)
	}

	wg.Wait()
}