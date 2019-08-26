package main

import "fmt"

func main() {
	var m map[string]string
	m = make(map[string]string)
	m["ex"] = "Prawat"
	fmt.Println(m["ex"])
}
