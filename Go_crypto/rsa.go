package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 生成 PrivateKey (一套)
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	fmt.Println("Key Set:", key)

	// 使用 RSA-OAEP 算法

	// 公钥 加密
	// 可能是 BUG or 我不懂 —— 用函数包装 public() 返回的是指针，而 publicKey 是原值
	// 一般用的是 PublicKey
	pubKey := key.Public()
	pubKey2 := key.PublicKey
	fmt.Printf("Public Key: %v %T\n", pubKey, pubKey)
	fmt.Printf("Public2 Key: %v %T\n", pubKey2, pubKey2)

	// 私钥 解密
	priKey := key.D
	fmt.Println("Private Key:", priKey)

	fmt.Println("加密")

	// 加密字段转为字节形式
	str := "Hello rsa."
	b := []byte(str)
	h := sha256.New() // 需要用到一个任意的 hash 器
	var label []byte  // label 是加密后串的明文，可为空
	label = []byte("greeting")
	res, err := rsa.EncryptOAEP(h, rand.Reader, &pubKey2, b, label)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	fmt.Println("解密")
	// 解密用到的 PubKey 是原来生成的一套 key， 同时 label 须相同
	ori, err := rsa.DecryptOAEP(h, rand.Reader, key, res, label)
	if err != nil {
		panic(err)
	}
	fmt.Println(ori)
	fmt.Println(string(ori))
}
