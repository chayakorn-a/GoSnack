package main

import "fmt"

func main() {
	num := 0
	for i := 0; i < 10; i++ {
		num += i
	}
	fmt.Println(num)

	num = 0
	i := 0
	// act like while loop
	for i < 10 {
		num += i
		i++
	}
	fmt.Println(num)

	names := []string{"game", "bank", "samui"}
	for i, name := range names {
		fmt.Println(name, i)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

}
