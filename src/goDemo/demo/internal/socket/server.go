package socket

import (
	"fmt"
	"io"
	_ "io"
	"net"
)

func process(conn net.Conn) {
	//关闭 conn
	defer conn.Close()

	//这里循环接收客户端发送的数据
	for {
		buf := make([]byte, 1024)
		fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())

		//1. 等待客户端通过 conn 发送信息
		//2. 如果客户端没有 wrtie[发送]，那么协程就阻塞在这里
		n, err := conn.Read(buf)
		if err == io.EOF || err != nil {
			fmt.Printf("客户端退出 err=%v", err)
			return
		}

		//3. 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func Server() {
	fmt.Println("服务器开始监听....")

	//1. tcp 表示使用的网络协议是 tcp
	//2. 0.0.0.0:8888 表示在本地监听 8888端口
	//3. 0.0.0.0 此方式支持 IPV4, IPV6。而 127.0.0.1 方式只支持 IPV4
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err= ", err)
		return
	}
	fmt.Printf("listen suc= %v\n", listen)

	//延时关闭 listen
	defer listen.Close()

	//循环等待客户端来链接我
	for {
		//等待客户端链接
		fmt.Println("等待客户端来链接....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
			continue
		} else {
			fmt.Printf("Accept() suc con= %v 客户端ip= %v\n", conn, conn.RemoteAddr().String())
		}
		//开启一个协程，为客户端服务
		go process(conn)
	}

}
