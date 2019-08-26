package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "Zero"
	a[1] = "One"
	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)

	b := [2]int{1, 2}
	fmt.Println(b)

	// ... make complier fix it on initial
	// below is as 6 once built
	c := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(c))
}
