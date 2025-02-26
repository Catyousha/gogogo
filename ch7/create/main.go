package main

import (
	"fmt"
	"time"
)

func main()  {
	go func (x int) {
		fmt.Println(x)
	}(10)
	time.Sleep(time.Second)
}