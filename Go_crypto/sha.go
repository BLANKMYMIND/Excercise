package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func main() {
	str := "how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you.how are you? I am fine, thank you."
	b := []byte(str)
	s := sha1.New()
	// new 之后可以先 reset 一下 s 的内容
	s.Reset()
	s.Write(b)
	// Sum 的参数是要加入结果中的前缀
	pre := []byte{0, 1, 2, 3, 4, 9, 9, 9, 9}
	res := s.Sum(nil)
	pres := s.Sum(pre)
	// 无论 b 字节串多长、多短，生成的 字节串总是长为 20.
	fmt.Println(b)
	fmt.Println(res)
	fmt.Println(pres)
	fmt.Println(pre)

	fmt.Println()

	// 256，512 用法都是相近的
	// 默认情况下，他们生成的 key 位数不同
	// sha1     20 位
	// sha224   28
	// sha256   32
	// sha348   48
	// sha512   64
	s = sha256.New()
	s.Reset()
	s.Write(b)
	fmt.Println(s.Sum(nil))
}
