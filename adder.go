package main

import "fmt"

// Closure function
// Get and call inside functon but still same local variable
// Use in case need to wrap some control state variable between function or each call
func adder() (func() int, func() int) {
	num := 0

	return func() int {
			num += 1
			return num
		},
		func() int {
			return num
		}
}

func main() {
	inc, cur := adder()

	fmt.Println(cur())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(cur())

	inc, cur = adder()

	fmt.Println(cur())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(cur())
}
