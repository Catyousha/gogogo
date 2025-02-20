package main

import (
	"encoding/json"
	"fmt"
)

type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"created"`
}

func main() {
	useall := UseAll{
		Name:    "Awanama",
		Surname: "Wijaya",
		Year:    2025,
	}

	// encode json object to str
	// Value {"username":"Awanama","surname":"Wijaya","created":2025}
	t, err := json.Marshal(&useall)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %s\n", t)
	}

	// decode from json string to object
	// Data type: main.UseAll with value {M. Ts 2020}
	str := `{"username": "M.", "surname": "Ts", "created":2020}`
	jsonRecord := []byte(str)
	temp := UseAll{}
	err = json.Unmarshal(jsonRecord, &temp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data type: %T with value %v\n", temp, temp)
	}
}
