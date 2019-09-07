package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	ID     string `json:id`
	Title  string `json:title`
	Status string `json:status`
}

func main() {
	t := Todo{
		ID:     "1",
		Title:  "Title",
		Status: "Testing",
	}

	b, err := json.Marshal(t)
	fmt.Printf("%T => %v \n %s \n", b, b, b)
	fmt.Println(err)
}
