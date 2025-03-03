package shared

import (
	"sync"
	"sync/atomic"

	"cty.sh/wc/pkg/common"
)

// CountAllConcurrent counts lines, words, and characters in a file concurrently using shared memory
func CountAllConcurrent(file string) (common.FileStats, error) {
	var lines, words, chars int64
	var wg sync.WaitGroup
	var err error
	var mu sync.Mutex

	// Count lines concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		lineCount, lineErr := common.CountLines(file)
		if lineErr != nil {
			mu.Lock()
			err = lineErr
			mu.Unlock()
			return
		}
		atomic.StoreInt64(&lines, int64(lineCount))
	}()

	// Count words concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		wordCount, wordErr := common.CountWords(file)
		if wordErr != nil {
			mu.Lock()
			err = wordErr
			mu.Unlock()
			return
		}
		atomic.StoreInt64(&words, int64(wordCount))
	}()

	// Count chars concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		charCount, charErr := common.CountChars(file)
		if charErr != nil {
			mu.Lock()
			err = charErr
			mu.Unlock()
			return
		}
		atomic.StoreInt64(&chars, int64(charCount))
	}()

	// Wait for all goroutines to complete
	wg.Wait()

	// Check if any errors occurred
	if err != nil {
		return common.FileStats{}, err
	}

	return common.FileStats{
		Lines: int(atomic.LoadInt64(&lines)),
		Words: int(atomic.LoadInt64(&words)),
		Chars: int(atomic.LoadInt64(&chars)),
		Path:  file,
	}, nil
}
