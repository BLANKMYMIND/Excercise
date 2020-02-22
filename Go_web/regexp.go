package main

import (
	"fmt"
	"regexp"
)

func main() {
	// regexp 是线程安全的，可以放心用
	// 使用 MustCompile 方法，在编译正则不成功时将会报 panic 来终止程序
	reg := regexp.MustCompile(`^1(?:(?:3[\d])|(?:4[5-7|9])|(?:5[0-3|5-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\d])|(?:9[1|8|9]))\d{8}$`)
	fmt.Print(reg.MatchString("15570877777"))
}
