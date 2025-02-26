package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			fmt.Println("Ended.")

		// listen timeout to prevent deadlock
		case <-time.After(4 * time.Second):
			fmt.Println("time.After()!")
			return
		}
	}
}

func main() {
	rand.NewSource(time.Now().Unix())

	createNumber := make(chan int)
	end := make(chan bool)

	n := 3

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		gen(0, 2*n, createNumber, end)
	}()

	for range n {
		fmt.Println("Created number:", <-createNumber)
	}

	end <- true
	wg.Wait()
}
