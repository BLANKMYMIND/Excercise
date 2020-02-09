package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
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
type Arity int

func (t *Arity) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arity) Divide(args *Args, quotient *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by 0")
	}
	quotient.Quo = args.A / args.B
	quotient.Rem = args.A % args.B
	return nil
}

func main() {
	// 新建类型的对象
	arity := new(Arity)
	// 注册类型的对象到 rpc
	err := rpc.Register(arity)
	if err != nil {
		fmt.Println(err)
	}

	// grpc 使用 http 协议
	// 注册 rpc 到 HTTP-Handle
	rpc.HandleHTTP()

	// 开启 http 监听
	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}
