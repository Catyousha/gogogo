package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

type PhoneBook []Entry

var MIN = 0
var MAX = 26
var CSVFILE = "./data/csv.data"

var data = PhoneBook{}
var index map[string]int

// exercise 1: Create a slice of structures using a structure that you created and sort the
// elements of the slice using a field from the structure
func (a PhoneBook) Len() int {
	return len(a)
}

func (a PhoneBook) Less(i, j int) bool {
	if a[i].Surname == a[j].Surname {
		return a[i].Name < a[j].Name
	}
	return a[i].Surname < a[j].Surname;
}

func (a PhoneBook) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func readCSVFile() error {
	f, err := os.OpenFile(CSVFILE, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := Entry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		}
		data = append(data, temp)
	}

	return nil
}

func saveCSVFile() error {
	f, err := os.OpenFile(CSVFILE, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	csvWriter := csv.NewWriter(f)
	for _, row := range data {
		_ = csvWriter.Write([]string{
			row.Name,
			row.Surname,
			row.Tel,
			row.LastAccess,
		})
	}
	csvWriter.Flush()

	return nil
}

func createIndex() error {
	index = make(map[string]int)
	for i, k := range data {
		index[k.Tel] = i
	}
	return nil
}

func initS(name, surname, tel string) (*Entry, error) {
	if tel == "" || surname == "" {
		return nil, fmt.Errorf("telephone or surname is empty")
	}

	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{
		Name:       name,
		Surname:    surname,
		Tel:        tel,
		LastAccess: LastAccess,
	}, nil
}

func insert(pS *Entry) error {
	// If it already exists, do not add it
	_, ok := index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}
	data = append(data, *pS)
	// Update the index
	_ = createIndex()

	err := saveCSVFile()
	if err != nil {
		return err
	}
	return nil
}

func deleteEntry(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found", key)
	}

	data = append(data[:i], data[i+1:]...)
	// Update the index - key does not exist any more
	delete(index, key)

	err := saveCSVFile()
	if err != nil {
		return err
	}
	return nil
}

func search(key string) *Entry {
	i, ok := index[key]
	if !ok {
		return nil
	}
	data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	_ = saveCSVFile()
	return &data[i]
}

func list(isReverse bool) {
	if isReverse {
		sort.Reverse(data)
	} else {
		sort.Sort(data)
	}
	for _, v := range data {
		fmt.Println(v)
	}
}

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`\d+$`)
	return re.Match(t)
}

func main() {
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
	defer f.Close()

	// set log's prefix (iLog yyyy/mm/dd HH:mm:ss <log text>)
	iLog := log.New(f, "iLog ", log.LstdFlags)

	iLog.Println("Initializing data...")
	err = readCSVFile()
	if err != nil {
		fmt.Println(err)
		iLog.Println("Err:", err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		iLog.Println("Err:", err)
		return
	}

	iLog.Println("Data has been initialized!")

	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list|insert|delete <args>\n", exe)
		iLog.Println("Required args not fulfilled.")
		return
	}

	switch arguments[1] {
	case "insert":
		if len(arguments) != 5 {
			fmt.Println("Usage: insert Name Surname Telephone")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		temp, err := initS(arguments[2], arguments[3], t)
		if err != nil {
			iLog.Println(err)
			fmt.Println(err);
			return
		}

		err = insert(temp)
		if err != nil {
			iLog.Println(err)
			fmt.Println(err);
			return
		}
		
	case "delete": 
		if len(arguments) != 3 {
			fmt.Println("Usage: delete Number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
			return
		}
		err := deleteEntry(t)
		if err != nil {
			fmt.Println(err)
		}

	case "search":
		iLog.Println("Searching for " + arguments[2] + "...")
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Not a valid telephone number:", t)
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
		list(false)
		iLog.Println("Listing all data.")

	// exercise 3: Add support for a reverse command to phonebook.go in order to list its
	// entries in reverse order
	case "reverse":
		list(true)
		iLog.Println("Listing all data in reverse.")
	default:
		fmt.Println("Not a valid option")
		iLog.Println("Unknown option.")
	}
}
