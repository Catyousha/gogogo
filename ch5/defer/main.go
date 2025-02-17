package main

import "fmt"

func d1() {
	for i := 3; i > 0; i-- {
		fmt.Print("D1-",i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print("D2-", i, " ")
		}()
	}
	fmt.Println("\nFrom D2")
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print("D3-", n, " ")
		} (i)
		fmt.Println("From", i, "iteration of D3")
	}
}

func main() {
	// order of exec:
	// println d2 -> defered d2 func
	// -> each iter println d3 -> defered d3 func
	// -> defered main line (d1())
	
	// From D2
	// D2-1 D2-2 D2-3 Hi
	// From 3 iteration of D3
	// From 2 iteration of D3
	// From 1 iteration of D3
	// D3-1 D3-2 D3-3 
	// Done
	// D1-3 D1-2 D1-1
	defer d1()
	d2()
	fmt.Println("Hi")
	d3()
	fmt.Println("\nDone")
}