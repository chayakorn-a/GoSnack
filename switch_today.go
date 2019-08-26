package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()
	fmt.Println(today)

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	default:
		fmt.Println("Faw Away")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Afternoon")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}
