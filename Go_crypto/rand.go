package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// 新建一个 big.Int 变量，其值为 255
	a := new(big.Int).SetUint64(uint64(255))

	// 生成 随机数，该随机数小于 a 的值 255 (0 ~ 255)
	c, err := rand.Int(rand.Reader, a)
	fmt.Println(c, err)

	// byte 其内是 uint8 所以直接赋值数字即可，数字可为十六进制
	key := []byte{48, 49, 0xff}
	fmt.Println(key)

	// 生成一个 16字节 (AES-128，128位) 的 key
	key = make([]byte, 16)
	for i, _ := range key {
		r, _ := rand.Int(rand.Reader, a)
		// fmt.Println(r)
		fmt.Println(r.Bytes()) // 生成的 bytes 会【自收缩】！(即 666 -> [2 154], 111 -> [111])
		key[i] = r.Bytes()[0]
	}
	fmt.Println(key)
}
