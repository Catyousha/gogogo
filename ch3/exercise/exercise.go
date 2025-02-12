package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// nameSurRE.go
func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

// intRE.go
func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}

// fieldsRE.go
func matchRecord(s string, d string) bool {
	fields := strings.Split(s, d)
	if len(fields) != 3 || !matchNameSur(fields[0]) || !matchNameSur(fields[1]) {
		return false
	}

	return matchTel(fields[2])
}

// csvData.go
var CSVFILE = "./data/data.csv"
func readCSVFile(comma string) ([][]string, error) {
	f, err := os.OpenFile(CSVFILE, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.Comma = rune(comma[0])
	lines, err := reader.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func main() {
	// exercise 1: existing array to map
	var arr = [3]int{123, 231, 312}
	voucher := make(map[string]int)
	for i, e := range arr {
		voucher[fmt.Sprintf("A%d", i)] = e
	}

	// exercise 2: converts existing map to two slices
	keys := make([]string, 0, len(voucher))
	values := make([]int, 0, len(voucher))
	for k, v := range voucher {
		keys = append(keys, k)
		values = append(values, v)
	}
	fmt.Println("Keys:", keys)
	fmt.Println("Values:", values)

	arguments := os.Args
	switch arguments[1] {
	// exercise 3:
	// nameSurRE.go to be able to process multiple command-line arguments
	case "check-surname":
		fmt.Println("CHECK SURNAME:")
		for i := 2; i < len(arguments); i++ {
			fmt.Printf("%s %t\n", arguments[i], matchNameSur(arguments[i]))
		}

	// exercise 4:
	// intRE.go to process multiple command-line arguments
	// and display totals of true and false results at the end
	case "check-int":
		fmt.Println("CHECK INT:")
		t := 0
		for i := 2; i < len(arguments); i++ {
			isTrue := matchInt(arguments[i])
			if isTrue {
				t = t + 1
			}
			fmt.Printf("%s %t\n", arguments[i], isTrue)
		}
		fmt.Printf("True: %d; False: %d;\n", t, (len(arguments)-2)-t)
	
	case "check-record":
		fmt.Println("CHECK RECORD:")
		if(len(arguments) < 4) {
			fmt.Println(matchRecord(arguments[2], ","))
		} else {
			fmt.Println(matchRecord(arguments[2], arguments[3]))
		}
	
	// exercise 5 & 8:
	// Make changes to csvData.go to separate the fields of a record
	// based on the character that is given as a command-line argument.
	case "check-csv":
		fmt.Println("CHECK CSV:")
		if(len(arguments) < 3) {
			fmt.Println(readCSVFile(","))
		} else {
			fmt.Println(readCSVFile(arguments[2]))
		}
	}

}
