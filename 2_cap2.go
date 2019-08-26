package main

func main() {
	// make([]int, len , cap)

	j := make([]int, 5)    // j len 5
	k := make([]int, 0, 5) // len 0 but cap 5, to use need to append only
	k = k[:cap(k)]         // extend the cap to maxinum
}
