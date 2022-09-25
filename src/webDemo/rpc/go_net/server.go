package go_net

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type World struct {
}

// SayHello 定义一个远程调用方法，该方法必须是公开的方法，且只能有两个可序列化的参数，且第二个参数必须是指针类型
//req 是客户端传递过来的数据，
//res 是返回给客户端的数据，
//两个参数类型不能是：channel 通道、complex 复数类型、func 函数，因为它们均不能序列化
func (world *World) SayHello(req string, res *string) error {
	*res = req + "hello"
	return nil
}

func RpcServerMain() {
	//注册 RPC 服务，name 是服务名称，any 是服务绑定的服务对象
	err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println(err)
	}

	//注册监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	for {
		fmt.Println("开始建立连接")

		//建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		//绑定服务
		rpc.ServeConn(conn)
	}
}

func RpcServerJson() {
	//注册 RPC 服务，name 是服务名称，any 是服务绑定的服务对象
	err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println(err)
	}

	//注册监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	for {
		fmt.Println("开始建立连接")

		//建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		//使用 jsonrpc 协议绑定服务
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
