package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{"String value", -12.123, Secret{
		"Mihalis",
		"Tsoukalos",
	}}

	v := reflect.ValueOf(true)
	// bool
	fmt.Println(v.Type().Name())

	r := reflect.ValueOf(A);

	// String value: <main.Record Value>
	fmt.Println("String value:", r.String());

	iType := r.Type()
	// i Type: main.Record
	fmt.Println("i Type:", iType)

	// Field1 string
	// {Field1, Field2, Field3} > Field1 > name , Field1 > string > name
	fmt.Println(r.Type().Field(0).Name, r.Field(0).Type().Name())

	// The 3 fields of main.Record are
    //     Field1  with type: string       and value _String value_
    //     Field2  with type: float64      and value _-12.123_
    //     Field3  with type: main.Secret  and value _{Mihalis Tsoukalos}_
	// Secret
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		fmt.Printf("\t%s ", iType.Field(i).Name);
		fmt.Printf("\twith type: %s ", field.Type())
		fmt.Printf("\tand value _%v_\n", field.Interface())
		
		k := reflect.TypeOf(field.Interface()).Kind()
		if k == reflect.Struct {
			fmt.Println(field.Type().Name())
		}
	}


}
