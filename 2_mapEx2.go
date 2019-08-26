package main

import "fmt"

type Student struct {
	name  string
	class string
	age   int
}

var m = map[string]Student{
	"000001": Student{"Chayakorn", "611", 12},
	"000002": Student{"OAT", "612", 30},
}

func main() {
	fmt.Println(m)
	delete(m, "000001")
	fmt.Println(m)

	fmt.Printf("% #v\n", m["000003"]) // not in the list then all fields are zero value

	// map return 2 parameter back
	if v, ok := m["000002"]; ok {
		fmt.Printf("found %s", v)
	} else {
		fmt.Printf("not found")
	}
}
