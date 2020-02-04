package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var tmp, res float64 = 1, 1
	for i := 0; tmp != x && i < 10; i++ {
		res -= (res*res - x) / (2 * res)
		tmp = res * res
	}
	return res
}

func main() {
	fmt.Println(Sqrt(2))
}
