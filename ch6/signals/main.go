package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) {
	fmt.Println("handleSignal() Caught:", sig)
}

func main() {
	fmt.Printf("Process ID: %d\n", os.Getpid())
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs)
	
	start := time.Now()
	go func () {
		for {
			// wait for data from sigs channel and store in sig var
			sig := <- sigs

			switch sig {
			// ctrl+c == SIGINT
			case syscall.SIGINT:
				duration := time.Since(start)
				fmt.Println("Execution time:", duration)
			
			// ctrl+t == SIGINFO
			// handleSignal() Caught: information request
			case syscall.SIGINFO:
				// do not use return here because the goroutine exits
				// but the time.Sleep() will continue to work!
				handleSignal(sig)
				os.Exit(0)

			default:
				fmt.Println("Caught:", sig)
			}
		}
	}()

	for {
		fmt.Print("+")
		time.Sleep(10 * time.Second)
	}
}