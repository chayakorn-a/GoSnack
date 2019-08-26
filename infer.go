package main

import "fmt"

const PI = 3.14
const (
	Zero = iota + 6
	One
	two
	three
	four
	five
)

type weekday int

const (
	Sunday weekday = iota
	Monday
	Tuesday
)

func main() {
	v := 42
	fmt.Printf("%T\n", v)

	const Hh = "Test"
	fmt.Printf("%T : %f \n%T : %s \n", PI, PI, Hh, Hh)

	fmt.Printf("%d,%d,%d,%d,%d,%d\n", Zero, One, two, three, four, five)

	fmt.Printf("%d %d %d", Sunday, Monday, Tuesday)
}
