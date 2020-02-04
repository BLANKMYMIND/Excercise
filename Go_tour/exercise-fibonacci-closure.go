package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	now := 0
	prev := 0
	return func() int {
		if now == 0 && prev == 0 {
			now = 1
			return 0
		} else {
			tmp := now
			now = now + prev
			prev = tmp
			return tmp
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
