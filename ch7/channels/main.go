package main

import (
	"fmt"
	"sync"
)

func writeOnlyCh(ch chan <- int, x int) {
	ch <- x
}

func readOnlyCh(ch <- chan int) int {
	return <- ch
}

func main() {
	c := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(c chan int) {
		defer wg.Done()
		writeOnlyCh(c, 10)
		close(c)

		fmt.Println("Exit.")
	}(c)

	// Exit.
	// Read: 10
	fmt.Println("Read:", readOnlyCh(c))
	
	// Is channel open: false
	_, ok := <- c
	fmt.Println("Is channel open:", ok);
	
	var ch chan bool = make(chan bool)
	
	// sender
	for i := 0; i < 5; i++ {
		go func(c chan bool) {
			c <- true
		}(ch)
	}


	// if close n < 5, will throw panic: send on closed channel
	n := 0
	
	for i := range ch {
		fmt.Println("i:", i)
		if i == true {
			n++
		}

		// can cause deadlock while iterating ch if not closed
		if n == 5 {
			fmt.Println("Closing at n:", n)
			close(ch)
			break
		}
	}

	// receiver
	for i := 0; i<5; i++ {
		fmt.Println(<-ch)
	}
}