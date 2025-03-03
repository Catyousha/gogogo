package buffered

import (
	"cty.sh/wc/pkg/common"
)

// CountAllConcurrent counts lines, words, and characters in a file concurrently using buffered channels
func CountAllConcurrent(file string) (common.FileStats, error) {
	// Create buffered channels for each count type
	linesChan := make(chan int, 1)
	wordsChan := make(chan int, 1)
	charsChan := make(chan int, 1)
	errorChan := make(chan error, 3)

	// Count lines concurrently
	go func() {
		lines, err := common.CountLines(file)
		if err != nil {
			errorChan <- err
			return
		}
		linesChan <- lines
	}()

	// Count words concurrently
	go func() {
		words, err := common.CountWords(file)
		if err != nil {
			errorChan <- err
			return
		}
		wordsChan <- words
	}()

	// Count chars concurrently
	go func() {
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
