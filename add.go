package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))

	fmt.Println(swap(1, 2))
}

func swap(i, j int) (int, int) {
	return j, i
}

func add(i, j int) int {
	return i + j
}
