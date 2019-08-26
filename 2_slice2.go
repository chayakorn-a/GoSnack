package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s, "Len:", len(s), "Capacity:", cap(s))
	p := []int(s)

	s = s[1:4]
	fmt.Println(s, "Len:", len(s), "Capacity:", cap(s))

	s = s[:2]
	fmt.Println(s, "Len:", len(s), "Capacity:", cap(s))

	s = s[1:]
	fmt.Println(s, "Len:", len(s), "Capacity:", cap(s))

	// Cap start from "starter of slice" until "end of original block"

	fmt.Println(p)
}
