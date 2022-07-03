package socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Client() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	defer conn.Close()

	//os.Stdin 代表标准输入，即终端键盘输入
	//在 win 或 linux 系统中所有输入输出都是通过文件中的内容进行传输实现的，所以当机器磁盘满时，也会出现网络阻塞
	reader := bufio.NewReader(os.Stdin)

	for {
		//从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err= ", err)
		}

		//如果用户输入的是 exit 就退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出..")
			break
		}

		//将 line 发送给服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
	}
}
