package main

import "fmt"

type person struct {
	name string
}

func (p person) Walk() {
	fmt.Println(p.name, "Walking")
}
func (p person) Eat() {
	fmt.Println(p.name, "eating")
}
func (p person) Greeting() {
	fmt.Println("Hello", p.name)
}
func (p *person) setter(ss string) {
	p.name = ss
}
func (p person) getter() string {
	return p.name
}

func main() {
	p := person{"test"}

	p.setter("OAT")
	p.Eat()
	p.Walk()
	p.Greeting()
	fmt.Println(p.getter())
}
