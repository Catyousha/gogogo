package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// start := time.Now()

	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
		return
	}
	dateString := os.Args[1]
	
	// go run . "31 December 2042"
	// Full: 2042-12-31 00:00:00 +0000 UTC
	// Time: 31 December 2042
	d, err := time.Parse("02 January 2006", dateString);
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	}

	// go run . "31 December 2042 21:03"
	// Full: 2042-12-31 21:03:00 +0000 UTC
	// Date: 31 December 2042
	// Time: 21 3
	d, err = time.Parse("02 January 2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	
}
