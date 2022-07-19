package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"goDemo/project/chatroom/internal/pkg/message"
	"net"
)

// Transfer 将这些方法关联到结构体中
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	//传输时，使用缓冲
	Buf [8096]byte
}

func (transfer *Transfer) ReadPkg() (mes message.Message, err error) {

	fmt.Println("读取客户端发送的数据...")

	//conn.Read 在 conn 没有被关闭的情况下，才会阻塞。如果客户端关闭了 conn，则就不会阻塞
	_, err = transfer.Conn.Read(transfer.Buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	//自定义通信协议，数据读取发送规则
	//根据 buf[:4] 转成一个 uint32 类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(transfer.Buf[0:4])

	//根据 pkgLen 读取消息内容
	n, err := transfer.Conn.Read(transfer.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}

	//把 pkgLen 反序列化成 -> message.Message
	err = json.Unmarshal(transfer.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}

func (transfer *Transfer) WritePkg(data []byte) (err error) {

	//自定义通信协议，数据读取发送规则
	//先发送一个长度给对方
	var pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(transfer.Buf[0:4], pkgLen)

	//发送长度
	n, err := transfer.Conn.Write(transfer.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//发送 data 本身
	n, err = transfer.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}
