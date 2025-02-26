package main

import (
	"fmt"
	"sync"
	"time"
)

type secret struct {
	RWM sync.RWMutex
	password string
}

var Password *secret
var wg sync.WaitGroup

func Change(pass string) {
	fmt.Println("Change() function")

	Password.RWM.Lock()
	fmt.Println("Change() locked")
	
	time.Sleep(4 * time.Second)
	Password.password = pass

	Password.RWM.Unlock()
	fmt.Println("Change() unlocked")
}


func show()  {
	defer wg.Done()

	Password.RWM.RLock()
	fmt.Println("Show function locked!")
	
	time.Sleep(2 * time.Second)
	
	fmt.Println("Pass value:", Password.password)
	defer Password.RWM.RUnlock()
}

func main() {
	Password = &secret{password: "myPass"}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go show()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("123456")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("654321")
	}()

	wg.Wait()
	
	// Change() function
	
	// Show function locked!
	// Show function locked!
	// Show function locked!

	// Change() function
	
	// Pass value: myPass
	// Pass value: myPass
	// Pass value: myPass
	
	/// race condition still occurs between two goroutine
	/// but it doesn't modify `secret.password` at the same time (no error when running with `-race` flag)
	// Change() locked
	// Change() unlocked
	
	// Change() locked
	// Change() unlocked
	
	// Current pass: 123456
	
	fmt.Println("Current pass:", Password.password)
}