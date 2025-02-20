// exercise 3: Create a full version of the wc(1) UNIX utility using commands instead of
// command-line options with the help of the cobra package.
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

// exercise 1: Use the functionality of byCharacter.go, byLine.go, and byWord.go in order
// to create a simplified version of the wc(1) UNIX utility.

func charByChar(file string) (n int, err error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("err reading file %s", err)
			break
		}
		n += len(string(line))
	}
	return;
}

func lineByLine(file string) (n int, err error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// exercise 8: The byLine.go utility uses ReadString('\n') to read the input file. Modify
	// the code to use Scanner (https://golang.org/pkg/bufio/#Scanner) for
	// reading.
	s := bufio.NewScanner(f)
	for s.Scan() {
		n++
	}
	if err := s.Err(); err != nil {
		fmt.Printf("err reading file %s", err)
	}
	return;
}

func wordByWords(file string) (n int, err error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// exercise 9: Similarly, byWord.go uses ReadString('\n') to read the input file—modify
	// the code to use Scanner instead.
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		r := regexp.MustCompile(`[^\\s]+`)
		words := r.FindAllString(line, -1)
		n += len(words)
	}

	if err := s.Err(); err != nil {
		fmt.Printf("err reading file %s", err)
	}
	return;
}




var rootCmd = &cobra.Command{
	Use:   "wc",
	Short: "A word counter application",
	Long:  `This is a word counter application.`,
	Run: func(cmd *cobra.Command, args []string) {
		l, _ := lineByLine(args[0])
		w, _ := wordByWords(args[0])
		c, _ := charByChar(args[0])
		fmt.Printf("\t%d\t%d\t%d\t%s\n", l, w, c, args[0])
	},
}


func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
}
