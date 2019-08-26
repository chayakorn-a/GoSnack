package main

import "fmt"

func main() {
	// a is main array
	// s is underlining array of a
	a := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%T => % #v\n", a, a)

	// [] call slice
	// a[1:4] the type is slice also it's not normal array
	var s []int = a[1:4] // Slice work like pointer
	fmt.Printf("%T => % #v\n", s, s)
	ss := []int(a[1:4])
	fmt.Printf("%T => % #v\n", ss, ss)
	fmt.Printf("% #v\n", a[1:4])

	// Slice concept is reference
	// slice is referenct to original array
	// Change on slice will change on original...it's same place
	s[0] = 11
	fmt.Printf("%T => % #v\n", a, a)
}
