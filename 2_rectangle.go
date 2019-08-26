package main

import "fmt"

type rectangle struct {
	x, y int
}

func area(rec rectangle) int {
	fmt.Println("rectangle function")
	return rec.x * rec.y
}

func (rec rectangle) area() int {
	fmt.Println("rectangle method")
	return rec.x * rec.y
}

func main() {
	rec := rectangle{3, 4}
	fmt.Println("Area is", area(rec))
	fmt.Println("Area is", rec.area())
}
