package main

import (
	"fmt"
	"reflect"
)

type T struct {
	F1 int
	F2 string
	F3 float64
}

func main()  {
	A := T{
		1, "f2", 3.0,
	}
	// {1 f2 3}
	fmt.Println(A)

	// String value: <main.T Value>
	r := reflect.ValueOf(&A).Elem()
	fmt.Println("String value:", r.String())

	// 0: F1 int = 1
	// 1: F2 string = f2
	// 2: F3 float64 = 3
	typeOfA := r.Type()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		tOfA := typeOfA.Field(i).Name
		fmt.Printf("%d: %s %s = %v\n", i, tOfA, f.Type(), f.Interface())
		
		fk := f.Type().Kind()
		if fk == reflect.Int {
			f.SetInt(-100)
		} else if fk == reflect.String {
			f.SetString("Changed!")
		}
	}

	// A: {-100 Changed! 3}
	fmt.Println("A:", A)

}