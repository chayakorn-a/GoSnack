package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	ss := s[1:3]
	fmt.Println(ss, "Len:", len(ss), "Capacity:", cap(ss))
	//fmt.Println(ss[2]) can not ref over current slice

	//scenario 1
	/*{
		ss = append(ss, 9, 8, 7)
		fmt.Println(ss[2])

		fmt.Printf("%T => % #v\n", s, s)
		fmt.Printf("%T => % #v\n", ss, ss)
	}*/

	//scenario 2
	{
		ss = append(ss, 9, 8, 7, 6)
		fmt.Println(ss[2])

		fmt.Printf("%T => % #v\n", s, s)
		fmt.Printf("%T => % #v\n", ss, ss)
	}
}
