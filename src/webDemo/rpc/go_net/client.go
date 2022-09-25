package go_net

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func RpcClientMain() {
	//建立与 RPC 服务端的连接
	conn, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	var res string
	err = conn.Call("hello.SayHello", "我是客户端", &res)
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取到远程服务返回的数据
	fmt.Println(res)
}

func RpcClientJson() {
	//建立与 RPC 服务端的连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//建立基于 json 编解码的 rpc 服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var res string
	err = client.Call("hello.SayHello", "我是客户端", &res)
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取到远程服务返回的数据
	fmt.Println(res)
}
