package main

import (
	"fmt"
	"log"
	"os"
	"path"
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
