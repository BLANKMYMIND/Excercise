package main

import (
	"fmt"
)

type ErrNegativeSqrt struct {
	value float64
}

func (err *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Can't require %v's Sqrt.", err.value)
}

func Sqrt(x float64) (float64, error) {
	var err error
	var res float64
	if x < 0 {
		err = &ErrNegativeSqrt{x}
		return res, err
	}
	var tmp float64
	tmp, res = 1, 1
	for i := 0; tmp != x && i < 10; i++ {
		res -= (res*res - x) / (2 * res)
		tmp = res * res
	}
	return res, err
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
