package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	pic := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8((x + y) / 2)
		}
		pic[y] = row
	}

	//pic := [][]uint8{}

	/*for y := 0; y < dy; y++ {
		row := []uint8{}
		for x := 0; x < dx; x++ {
			row = append(row, uint8((x+y)/2))
		}
		pic = append(pic, row)
	}*/

	return pic
}

func main() {
	pic.Show(Pic)
}
