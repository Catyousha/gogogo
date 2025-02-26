package main

import (
	"fmt"
	"time"
)

func main()  {
	// Multiple goroutines:
	// 4
	// 2
	// 3
	// 0
	// 1
	fmt.Println("Multiple goroutines:");
	for i := 0; i < 5; i++ {
		go func(x int) {
			fmt.Println(x)
		}(i)
	}
	time.Sleep(time.Second);
}