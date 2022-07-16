package unitTest

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
)

func handleMsg(conn net.Conn) {
	fmt.Println("等待客户端发送消息...")
	defer conn.Close()

	bytes := make([]byte, 1024)
	for {
		n, err := conn.Read(bytes)
		if err != nil {
			fmt.Errorf("与客户端连接异常 %v\n", err)
			return
		}
		fmt.Printf("接收到消息：%v\n", string(bytes[:n]))
	}
}

func TestServer(t *testing.T) {
	t.Log("服务器开启...")
	t.Log("开始创建连接...")

	listen, err := net.Listen("tcp", "0.0.0.0:8890")
	if err != nil {
		t.Errorf("建立连接失败 %v\n", err)
	}
	defer listen.Close()

	for {
		t.Log("等待与客户端建立连接...")

		conn, err := listen.Accept()
		if err != nil {
			t.Errorf("与客户端连接失败 %v\n", err)
			continue
		}

		t.Logf("连接成功 %v\n", conn.RemoteAddr().String())
		go handleMsg(conn)
	}
}

func TestClient(t *testing.T) {
	t.Log("客户端开启...")
	t.Log("开始创建连接...")

	conn, err := net.Dial("tcp", "0.0.0.0:8890")
	if err != nil {
		t.Errorf("建立连接失败 %v\n", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		//此段代码没问题，只是不能在本 test 环境下阻塞等待输入
		value, err := reader.ReadString('\n')
		if err != nil {
			t.Errorf("客户端接收输入异常 %v\n", err)
			return
		}

		value = strings.Trim(value, " \r\n")
		if value == "END" {
			t.Log("结束通话...")
			return
		}
		conn.Write([]byte(value))
	}
}
