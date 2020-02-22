package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"math/big"
)

// 生成一个 16字节 (AES-128，128位) 的 key (密钥)
func newKey() []byte {
	a := new(big.Int).SetUint64(uint64(255))

	key := make([]byte, 16)
	for i, _ := range key {
		r, _ := rand.Int(rand.Reader, a)
		key[i] = r.Bytes()[0]
	}
	fmt.Println(key)
	return key
}

func main() {
	// 获取 key
	b := newKey()

	// 生成 Cipher (使用 key b 的 aes 加密解密器)
	c, err := aes.NewCipher(b)
	if err != nil {
		panic(err)
	}

	// 原字符串
	str := "Hello boys."

	// 字符串转为字节表示
	b1 := []byte(str)

	fmt.Println("Origin b1", b1)
	fmt.Println("b1 length", len(b1))

	// 注意！无论是加解密的字节长度都应该为 BlockSize 的整数倍！长度不够填充 0
	// 填充原始字节
	l := len(b1)
	if l%aes.BlockSize != 0 {
		l = ((l / aes.BlockSize) + 1) * aes.BlockSize
		b1 = b1[:l]
		fmt.Println("Origin new b1", b1)
		fmt.Println("b1 new length", len(b1))
	}

	// 按照原始字节的填充长度创建 b2, b3
	var b2 = make([]byte, l)
	var b3 = make([]byte, l)

	// 加密 b1 到 b2
	c.Encrypt(b2, b1)
	fmt.Println("Encrypt", b2)

	// 解密 b2 到 b3
	c.Decrypt(b3, b2)
	fmt.Println("Decrypt", b3)

	// 将 b3 转为字符串表示 (0 的填充不影响正常转义)
	s := string(b3)
	fmt.Println(s)
}
