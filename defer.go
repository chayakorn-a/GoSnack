package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("End")
		// we would use this tactic to put defer function at start of program
		// it will be last
	}() // () it mean this function is already called

	defer fmt.Println("World") // defer will be proceeded at the end of function
	fmt.Println("Hello")
	Test()

	// defer put in Stack then FILO
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func Test() {
	defer fmt.Println("In function")
}
