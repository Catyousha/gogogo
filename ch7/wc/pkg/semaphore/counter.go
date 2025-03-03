package semaphore

import (
	"context"

	"cty.sh/wc/pkg/common"
	"golang.org/x/sync/semaphore"
)

// CountAllConcurrent counts lines, words, and characters in a file concurrently using semaphores
func CountAllConcurrent(file string) (common.FileStats, error) {
	// Create a semaphore with weight 3 (max 3 concurrent operations)
	sem := semaphore.NewWeighted(3)
	ctx := context.Background()

	// Create channels for results
	linesChan := make(chan int, 1)
	wordsChan := make(chan int, 1)
	charsChan := make(chan int, 1)
	errorChan := make(chan error, 3)

	// Count lines concurrently
	go func() {
		// Acquire semaphore
		if err := sem.Acquire(ctx, 1); err != nil {
			errorChan <- err
			return
		}
		defer sem.Release(1)

		lines, err := common.CountLines(file)
		if err != nil {
			errorChan <- err
			return
		}
		linesChan <- lines
	}()

	// Count words concurrently
	go func() {
		// Acquire semaphore
		if err := sem.Acquire(ctx, 1); err != nil {
			errorChan <- err
			return
		}
		defer sem.Release(1)

		words, err := common.CountWords(file)
		if err != nil {
			errorChan <- err
			return
		}
		wordsChan <- words
	}()

	// Count chars concurrently
	go func() {
		// Acquire semaphore
		if err := sem.Acquire(ctx, 1); err != nil {
			errorChan <- err
			return
		}
		defer sem.Release(1)

		chars, err := common.CountChars(file)
		if err != nil {
			errorChan <- err
			return
		}
		charsChan <- chars
	}()

	// Collect results
	var stats common.FileStats
	stats.Path = file

	// Check for errors first (non-blocking)
	select {
	case err := <-errorChan:
		return common.FileStats{}, err
	default:
		// No errors, continue
	}

	// Collect results from all goroutines
	stats.Lines = <-linesChan
	stats.Words = <-wordsChan
	stats.Chars = <-charsChan

	return stats, nil
}
