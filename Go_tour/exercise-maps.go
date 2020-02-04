package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	strmap := strings.Fields(s)
	res := make(map[string]int)
	for i := range strmap {
		word := strmap[i]
		_, ok := res[word]
		if ok == true {
			res[word] += 1
		} else {
			res[word] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
