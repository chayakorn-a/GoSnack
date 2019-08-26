package main

import "fmt"

func compute(i int, j int, fn func(int, int) int) int {
	return fn(i, j)
}

func main() {
	TestFunc := func(i, j int) int {
		return i + j
	}

	fmt.Println(TestFunc(1, 2))
	fmt.Println(compute(1, 2, TestFunc))
}
