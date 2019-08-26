package main

import "fmt"

func main() {
	i := 42
	fmt.Println(i)

	var p *int
	fmt.Println(p)

	p = &i
	fmt.Println(p)

	fmt.Printf("%p %p\n", p, &i)

	*p = 33
	fmt.Println(i)
}
