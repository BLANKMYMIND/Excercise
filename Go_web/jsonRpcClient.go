package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/rpc/jsonrpc"
)

// CODE ACCORDING TO BOOK 8.4

// 参数数据 及 返回数据 都得使用结构体包装
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

// 包装用于注册的类型，底层什么类型都可以
// Arity 类型 8 用写
// type Arity int

func main() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{9, 7}
	var reply int
	err = client.Call("Arity.Multiply", args, &reply)
	if err != nil {
		log.Fatal("calling:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}
