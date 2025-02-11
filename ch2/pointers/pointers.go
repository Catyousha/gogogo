package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x + *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main()  {
	// Memory address of f: 0x14000112008
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)

	// Memory address of f: 0x14000112008
	fp := &f
	fmt.Println("Memory address of f:", fp)

	// Updated value of f: 2
	*fp = 2;
	fmt.Println("Updated value of f:", f)

	// Updated value of f through function: 4 (2 + 2)
	processPointer(fp)
	fmt.Println("Updated value of f through function:", f);

	// Value of x: 8
	// Unchanged value of f: 4
	x := returnPointer(f)
	fmt.Println("Value of x:", *x)
	fmt.Println("Unchanged value of f:", f)

	// Value of xx: 8
	// Unchanged value of fp: 4
	xx := bothPointers(fp);
	fmt.Println("Value of xx:", *xx);
	fmt.Println("Unchanged value of fp:", *fp)

	// Empty k: <nil> (point nowhere)
	var k *aStructure
	fmt.Println("Empty k:", k);

	// Value of k &{(0+0i) 0}
	// Default fields of k: (0+0i) 0
	k = new(aStructure);
	fmt.Println("Value of k", k)
	fmt.Println("Default fields of k:", k.field1, k.field2)
}