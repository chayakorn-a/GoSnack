package main

import "fmt"

func fibonaci() func() int {
	value1 := 0
	value2 := 0
	return func() int {
		if value2 == 0 {
			value2++
		} else {
			result := value1 + value2
			value1 = value2
			value2 = result
		}
		return value2
	}
}

func main() {
	f := fibonaci()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", f())
	}
}
