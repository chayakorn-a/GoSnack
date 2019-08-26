package main

import "fmt"

var c, python, java bool
var i int = 10

func main() {
	var i int = 5
	var ss = "name"
	sss := "เทส"

	fmt.Println(i, c, python, java, ss, sss)
	fmt.Printf("%T\n", i)

	var id float64 = float64(i)
	fmt.Printf("%T\n", id)
}
