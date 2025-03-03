package common

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

// FileStats represents the statistics of a file
type FileStats struct {
	Lines int
	Words int
	Chars int
	Path  string
}

// CountLines counts the number of lines in a file
func CountLines(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		n++
	}

	if err := s.Err(); err != nil {
		return 0, err
	}
	return n, nil
}

// CountWords counts the number of words in a file
func CountWords(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		r := regexp.MustCompile(`\S+`)
		words := r.FindAllString(line, -1)
		n += len(words)
	}

	if err := s.Err(); err != nil {
		return 0, err
	}
	return n, nil
}

// CountChars counts the number of characters in a file
func CountChars(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n := 0
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}
		n += len(line)
	}
	return n, nil
}

// CountAll counts lines, words, and characters in a file
func CountAll(file string) (FileStats, error) {
	lines, err := CountLines(file)
	if err != nil {
		return FileStats{}, err
	}

	words, err := CountWords(file)
	if err != nil {
		return FileStats{}, err
	}

	chars, err := CountChars(file)
	if err != nil {
		return FileStats{}, err
	}

	return FileStats{
		Lines: lines,
		Words: words,
		Chars: chars,
		Path:  file,
	}, nil
}
