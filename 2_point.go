package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// Same thing, but as a method of Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

type ColorPoint struct {
	Point // Point data type will be embeded to ColorPoint, don't declare the name
	// if declare name, it will be just data member, not embeded
	Color string
}

func main() {
	p := Point{1, 1}

	fmt.Println(Distance(p, Point{4, 4}))
	fmt.Println(p.Distance(Point{4, 4}))

	c := ColorPoint{p, "blue"} // p Point will be embeded to c
	fmt.Println(c.x)
	fmt.Println(c.Distance(Point{4, 4}))
}
