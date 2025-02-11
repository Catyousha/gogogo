package main

import (
	"errors"
	"fmt"
	"os"
)

var customMsg = "this is a custom error message"

func check(a complex64, b float32) error {
	if a == 0 {
		return errors.New(customMsg)
	}

	// the cooler error
	if b == 0 {
		return fmt.Errorf("a %e and b %e. UserID: %d", a, b, os.Getuid())
	}

	c1 := a + 1i;
	// cast float32 to complex64 (b + 1i)
	c2 := complex(b, 1);

	// addition of two complex64
	c3 := complex64(c1 + c2);

	// c1: (2.000000e+00+1.000000e+00i) (complex64)
	// c2: (3.000000+1.000000i) (complex64)
	// c3: (5.000000+2.000000i) (complex64)
	fmt.Printf("c1: %f (%T)\n", c1, c1)
	fmt.Printf("c2: %f (%T)\n", c2, c2)
	fmt.Printf("c3: %f (%T)\n", c3, c3);

	// cZero: (0.000000+0.000000i) (complex64)
	cZero := c3 - c3;
	fmt.Printf("cZero: %f (%T)\n", cZero, cZero)

	// div: (1.666667+0.666667i) (complex64)
	div := c3 / 3 // complex64 / int = int
	fmt.Printf("div: %f (%T) \n", div, div);

	return nil
}

func main() {
	err := check(0, 10)

	if err == nil {
		fmt.Println("check() ended normally")
	} else {
		// fmt.PrintLn(errors.New("...") || fmt.Errorf("..."))
		fmt.Println(err)
	}

	err = check(0, 0)
	if err.Error() == customMsg {
		fmt.Println("Custom error detected!")
	}

	check(2, 3)

}
