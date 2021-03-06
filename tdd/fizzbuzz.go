package tdd

import "strconv"

func FizzBuzz(num int) string {
	switch {
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(num)
}

/* Real code from Teacher
func FizzBuzz(num int) string {
 if isFizzBuzz(num) {
  return "FizzBuzz"
 }

 if isBuzz(num) {
  return "Buzz"
 }

 if isFizz(num) {
  return "Fizz"
 }

 return strconv.Itoa(num)
}

func isFizzBuzz(num int) bool {
 return num%15 == 0
}
func isBuzz(num int) bool {
 return num%5 == 0
}
func isFizz(num int) bool {
 return num%3 == 0
}*/
