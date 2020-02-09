package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
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
	// 新建对象并注册
	arity := new(Arity)
	rpc.Register(arity)

	// json-rpc 直接使用 tcp 协议 (注意是 tcp Addr 不是 ip Addr)
	// 生成地址结构
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8888")
	checkError(err)

	// 监听端口
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// 连接并转发
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// 直接将连接推给 jsonRpc 处理
		go jsonrpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		panic(err)
	}
}
