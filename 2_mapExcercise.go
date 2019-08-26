package main

import "fmt"

type Student struct {
	name  string
	class string
	age   int
}

func main() {
	var m map[string]Student
	m = make(map[string]Student)

	ID1 := "000001"
	// Key is student ID
	// Value is student Info: name, class, age
	m[ID1] = Student{"Chayakorn", "611", 12}
	fmt.Println(ID1, m[ID1].name, m[ID1].class, m[ID1].age)
}
