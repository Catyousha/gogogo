package main

import (
	"fmt"
	"runtime/metrics"
	"sync"
	"time"
)

func main() {
	// This metric name is used to track the number of active goroutines in the program
	// It's a built-in runtime metric provided by Go
	const nGo = "/sched/goroutines:goroutines"

	// Initialize a slice to hold one metric sample
	// We only need one sample since we're tracking a single metric (goroutine count)
	getMetric := make([]metrics.Sample, 1)
	getMetric[0].Name = nGo

	// Create 3 goroutines to demonstrate concurrent execution
	// Each goroutine sleeps for 4 seconds to simulate work
	// WaitGroup is used to wait for all goroutines to complete
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(4 * time.Second)
		}()
	}

	// Read the current metric value
	// This will show 4 goroutines: main + 3 created goroutines
	metrics.Read(getMetric)
	if getMetric[0].Value.Kind() == metrics.KindBad {
		fmt.Printf("metric %q no longer supported\n", nGo)
	}

	mVal := getMetric[0].Value.Uint64()
	fmt.Printf("Number of goroutines: %d\n", mVal)

	// Wait for all goroutines to finish
	// Then read metrics again to show the final count (should be 1, just main)
	wg.Wait()
	metrics.Read(getMetric)
	mVal = getMetric[0].Value.Uint64()
	fmt.Printf("Number of goroutines before exiting: %d\n", mVal)
}
