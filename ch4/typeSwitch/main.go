package main

import "fmt"

type Secret struct {
	SecretValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

func Teststruct(x interface{}) {
	switch T := x.(type) {
	case Secret:
		fmt.Println("secret")

	case Entry:
		fmt.Println("entry")

	default:
		fmt.Printf("unsupported type: %T\n", T)
	}
}

func Learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf("Data type: %T\n", T)
	}
}

func main() {
	A := Entry{100, "F2", Secret{"myPassword"}}

	// entry
	// secret
	// unsupported type: string
	Teststruct(A)
	Teststruct(A.F3)
	Teststruct("A string")

	// Data type: float64
	// Data type: int32
	Learn(12.23)
	Learn('€')
}
