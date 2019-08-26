package main

import (
	"fmt"
)

const V, I, X, L, C, D, M = "V", "I", "X", "L", "C", "D", "M"

// Logic to print out Roman for Single digits
func PrintRomanDigit(iDigit int, low, mid, high string) {
	switch {
	case iDigit <= 3:
		for i := 0; i < iDigit; i++ {
			fmt.Printf("%s", low)
		}
	case iDigit > 3 && iDigit <= 8:
		if iDigit < 5 {
			fmt.Printf("%s", low)
		}
		fmt.Printf("%s", mid)
		for i := 0; i < iDigit-5; i++ {
			fmt.Printf("%s", low)
		}
	default:
		if iDigit < 10 {
			fmt.Printf("%s", low)
		}
		fmt.Printf("%s", high)
	}
}

// Support to printout for maximum 3 digits (1-999)
func PrintRoman(iNum int) {
	var iDigit int

	iDigit = (iNum / 100) % 10 // make a digit of hundred
	if iDigit > 0 {
		PrintRomanDigit(iDigit, C, D, M)
	}
	iDigit = (iNum / 10) % 10 // make a digit of ten
	if iDigit > 0 {
		PrintRomanDigit(iDigit, X, L, C)
	}

	iDigit = iNum % 10 // make a digit of last digit
	if iDigit > 0 {
		PrintRomanDigit(iDigit, I, V, X)
	}
}

func main() {
	// Loop to print Roman number from 1 to 150
	for i := 0; i < 151; i++ {
		PrintRoman(i)
		fmt.Printf(" ")
		if i%10 == 0 {
			fmt.Printf("\n")
		}
	}
}
