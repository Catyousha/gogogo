package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// f1 demonstrates context cancellation using WithCancel
// It will complete if either:
// - The timer of t seconds expires, printing the time
// - The context is cancelled after 4 seconds, printing "Done"
// Since t=3 and cancel happens at 4s, the timer will complete first
func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1() Done:", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1():", r)
	}
	return
}

// f2 demonstrates context timeout using WithTimeout
// It will complete if either:
// - The context times out after t seconds
// - The timer of t seconds expires
// - The context is cancelled after 4 seconds
// Since timeout and timer are both t=3s, they will race,
// but both complete before the 4s cancel
func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2() Done:", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}
	return
}

// f3 demonstrates context deadline using WithDeadline
// It will complete if either:
// - The deadline expires after 2*t seconds
// - The timer of t seconds expires
// - The context is cancelled after 4 seconds
// Since t=3, the timer will complete at 3s, before both
// the 6s deadline and 4s cancel
func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3() Done:", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}
	return
}

func main() {
	t := 3
	fmt.Println("Delay:", t)
	// f3(): 2025-02-28 14:16:55.194774833 +0700 WIB m=+3.000550042
	// f2(): 2025-02-28 14:16:55.194821084 +0700 WIB m=+3.000595752
	// f1(): 2025-02-28 14:16:55.194631666 +0700 WIB m=+3.000406584

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done();
		f3(t)
	}()
	go func() {
		defer wg.Done();
		f2(t)
	}()
	go func() {
		defer wg.Done();
		f1(t)
	}()

	wg.Wait()
}
