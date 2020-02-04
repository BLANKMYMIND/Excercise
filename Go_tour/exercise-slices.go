package main

import "golang.org/x/tour/pic"

// import "math"

func Pic(dx, dy int) [][]uint8 {
	var all = make([][]uint8, dy)
	for y := range all {
		all[y] = make([]uint8, dx)
		for x := range all[y] {
			all[y][x] = uint8(x % (y + 1))
		}
	}
	return all
}

func main() {
	pic.Show(Pic)
}
