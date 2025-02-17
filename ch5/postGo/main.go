package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mactsouk/post05"
)

var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}

func main() {
	post05.Hostname = "localhost"
	post05.Port = 5432
	post05.Username = "user"
	post05.Password = "pass"
	post05.Database = "post05"

	data, err := post05.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range data {
		fmt.Println(v)
	}

	SEED := time.Now().Unix()
	rand.NewSource(SEED)
	random_username := getString(5)

	t := post05.Userdata{
		Username:    random_username,
		Name:        "Awanama",
		Surname:     "Wijaya",
		Description: "awaawawa",
	}

	id := post05.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	err = post05.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	// must be error
	err = post05.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	id = post05.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	t = post05.Userdata{
		Username:    random_username,
		Name:        "Wawanama",
		Surname:     "Wijayaya",
		Description: "This might not be me!"}

	err = post05.UpdateUser(t)
	if err != nil {
		fmt.Println(err)
	}
}
