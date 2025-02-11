package main

import "fmt"

func main()  {
	r := '€'
    // int32: 8364
    fmt.Println("int32:", r)
    // string: %!s(int32=8364), character: €
    fmt.Printf("string: %s, character: %c\n", r, r);

    aString := "Hello World! €"
    // %!s(int32=72) %!s(int32=101) %!s(int32=108)...
    for _, v := range aString {
        fmt.Printf("%s ", v)
    }

    // H e l l o   W o r l d !   €
    for _, v := range aString {
        fmt.Printf("%c ", v)
    }

    fmt.Printf("\n")
}