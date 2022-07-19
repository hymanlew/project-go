package main

import (
	"fmt"
	"goDemo/project/chatroom/internal/pkg/utils"
	"goDemo/project/chatroom/internal/server/process"
	"net"
	"time"
)

func init() {
	//当服务器启动时，就初始化 redis 连接池
	utils.InitPool("localhost:6379", 16, 0, 30*time.Second)
}

func main() {
	fmt.Println("服务器开始 8889 端口监听....")

	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	for {
		fmt.Println("等待客户端来链接服务器.....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
			continue
		}

		//一旦链接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}

//处理和客户端的通讯
func process(conn net.Conn) {

	//这里需要延时关闭conn
	defer conn.Close()

	//调用总控, 处理消息
	processor := serverProcess.Processor{
		Conn: conn,
	}
	err := processor.ServerProcess()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误=err", err)
		return
	}
}
