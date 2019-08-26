package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	AllWord := strings.Fields(s)
	len := len(AllWord)
	ret := map[string]int{}

	for i := 0; i < len; i++ {
		AllWord[i] = strings.Trim(AllWord[i], ".")
		AllWord[i] = strings.Trim(AllWord[i], ",")

		if strings.Count(s, AllWord[i]) > 1 {
			ret[AllWord[i]] = strings.Count(s, AllWord[i])
		}
	}

	return ret
}

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."

	fmt.Println(WordCount(s))
}
