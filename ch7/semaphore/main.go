package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

// Workers defines the maximum number of concurrent goroutines
var Workers = 4

// sem is a weighted semaphore that limits concurrent access to a resource
// Unlike WaitGroup which just waits for all goroutines to complete,
// semaphore controls how many goroutines can execute simultaneously
var sem = semaphore.NewWeighted(int64(Workers))

// worker simulates a task that takes time to complete
// It calculates the square of a number and waits for 1 second
func worker(n int) int {
	square := n * n
	time.Sleep(time.Second)
	return square
}

func main() {
	// Total number of jobs to process
	nJobs := 5

	// Slice to store results from all workers
	results := make([]int, nJobs)
	// Context for managing semaphore operations
	ctx := context.TODO()

	// Expected output (order may vary due to concurrent execution):
	// i: 0 r: 0
	// i: 1 r: 1
	// i: 3 r: 9
	// i: 2 r: 4
	// i: 4 r: 16

	// Process each job with controlled concurrency
	for i := range results {
		// Acquire 1 unit from semaphore before starting a new goroutine
		// This blocks if all Workers are busy, limiting concurrency
		// WaitGroup doesn't limit concurrency - it would start all goroutines at once
		err := sem.Acquire(ctx, 1)
		if err != nil {
			fmt.Println("Cannot acquire semaphore:", err)
			break
		}

		go func(i int) {
			// Release semaphore when goroutine completes
			// Similar to WaitGroup.Done() but also frees a resource slot
			defer sem.Release(1)
			r := worker(i)
			results[i] = r
			fmt.Println("i:", i, "r:", r)
		}(i)
	}

	// Wait for all workers to complete by trying to acquire all semaphore units
	// This is conceptually similar to WaitGroup.Wait() but with different mechanics
	// It blocks until all goroutines have released their semaphore units
	err := sem.Acquire(ctx, int64(Workers))
	if err != nil {
		fmt.Println(err)
	}

	// Print final results in order
	// 0 -> 0
	// 1 -> 1
	// 2 -> 4
	// 3 -> 9
	// 4 -> 16
	for k, v := range results {
		fmt.Println(k, "->", v)
	}
}
