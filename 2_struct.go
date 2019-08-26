package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	var xx Vertex
	xx.x = 1
	xx.y = 2
	fmt.Println(xx.x, xx.y)

	fmt.Println("+++Block2")
	v := Vertex{3, 4}
	fmt.Printf("% #v\n", v)
	v2 := Vertex{y: 9, x: 8}
	fmt.Printf("% #v\n", v2)

	fmt.Println("+++Block3")
	pV := &Vertex{3, 4}
	fmt.Printf("% #v\n", *pV)
	pV.x = 9
	fmt.Printf("% #v\n", *pV)

	fmt.Println("+++Block3")
	p := &v
	p.y = 16
	fmt.Printf("% #v\n", v)

}
