package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}

type Result struct {
	job    Client
	square int
}

var size = runtime.GOMAXPROCS(0)
var clients = make(chan Client, size)
var data = make(chan Result, size)

func worker(wg *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Result{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	wg.Done()
}

func create(n int) {
	defer close(clients)
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
}

func main() {
	nJobs := 10
	nWorkers := 5

	go create(nJobs)

	finished := make(chan interface{})
	sum := 0;
	go func() {
		// listen to data channels
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
			sum += d.square
		}
		finished <- true
	}()

	var wg sync.WaitGroup
	for range nWorkers {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(data)

	fmt.Printf("Finished: %v\n", <-finished)
	fmt.Println("Result:", sum)
}
