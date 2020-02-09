package main

import (
	"fmt"
	"net/rpc"
	"os"
)

// CODE ACCORDING TO BOOK 8.4

// 参数数据 及 返回数据 都得使用结构体包装
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

// Arity 类型 8 用加

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]

	// --------
	// 接通指定接口 http
	client, err := rpc.DialHTTP("tcp", serverAddress+":8888")
	defer client.Close()
	if err != nil {
		panic(err)
	}

	// 为参数结构体赋值
	args := Args{17, 8}
	// 声明返回结构体 / 类型
	var reply int
	// 调用函数
	// 注意 类型.函数 的 类型 不能缺少！
	err = client.Call("Arity.Multiply", args, &reply)
	// 一定要处理 err
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

}
