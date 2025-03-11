package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Result struct {
	Duration time.Duration
	Error    error
}

func main() {
	n := flag.Int("n", 100, "Number of requests to make")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL")
		return
	}

	url := os.Args[1]

	results := make(chan Result, *n)
	start := time.Now()

	// Create worker pool
	for i := 0; i < *n; i++ {
		go func() {
			reqStart := time.Now()
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			results <- Result{
				Duration: time.Since(reqStart),
				Error:    err,
			}
		}()
	}

	// Collect results
	var totalTime time.Duration
	var errors int
	var maxTime time.Duration
	var minTime = time.Hour

	for range *n {
		result := <-results
		if result.Error != nil {
			errors++
			continue
		}
		totalTime += result.Duration
		if result.Duration > maxTime {
			maxTime = result.Duration
		}
		if result.Duration < minTime {
			minTime = result.Duration
		}
	}

	totalDuration := time.Since(start)
	successfulRequests := *n - errors

	// Print results
	fmt.Printf("\nBenchmark Results for %s\n", url)
	fmt.Printf("Complete requests: %d\n", *n)
	fmt.Printf("Failed requests: %d\n", errors)
	fmt.Printf("Time taken for tests: %.2f seconds\n", totalDuration.Seconds())
	fmt.Printf("Requests per second: %.2f\n", float64(successfulRequests)/totalDuration.Seconds())
	if successfulRequests > 0 {
		fmt.Printf("Mean time per request: %.2f ms\n", totalTime.Seconds()*1000/float64(successfulRequests))
		fmt.Printf("Min time per request: %.2f ms\n", minTime.Seconds()*1000)
		fmt.Printf("Max time per request: %.2f ms\n", maxTime.Seconds()*1000)
	}
}
