package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	"golang.org/x/exp/rand"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{
	{
		"Mihalis", "Tsoukalos", "123",
	},
	{
		"Mary", "Jane", "213",
	},
	{
		"John", "Doe", "321",
	},
}
var MIN = 0
var MAX = 26

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}
	return temp
}


func populate(n int) {
	for i := 0; i < n; i++ {
		name := getString(4);
		surname := getString(5);
		n := strconv.Itoa(random(100, 199));
		data = append(data, Entry{name, surname, n})
	}
}

func main() {
	fmt.Println(os.TempDir())
	LOGFILE := path.Join(os.TempDir(), "phonebook.log")
	// O_APPEND = append new data at the end of file
	// O_CREATE = or, create when not exist
	// O_WRONLY = and set permission for writing only
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// execute when f variable unused anymore
	defer f.Close();

	// set log's prefix (iLog yyyy/mm/dd HH:mm:ss <log text>)
	iLog := log.New(f, "iLog ", log.LstdFlags);

	iLog.Println("Initializing data...")
	populate(100);
	iLog.Println("Data has been initialized!")
	
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <args>\n", exe)
		iLog.Println("Required args not fulfilled.")
		return
	}

	switch arguments[1] {
	case "search":
		iLog.Println("Searching for "+arguments[2]+"...")
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			iLog.Println(arguments[2], "not found.")
			return
		}
		iLog.Println(arguments[2], "found.")
		fmt.Println(*result)
	case "list":
		list()
		iLog.Println("Listing all data.")
	default:
		fmt.Println("Not a valid option")
		iLog.Println("Unknown option.")
	}
}
