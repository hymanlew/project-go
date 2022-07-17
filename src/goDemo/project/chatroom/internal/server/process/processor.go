package serverProcess

import (
	"fmt"
	"goDemo/project/chatroom/internal/pkg/message"
	"goDemo/project/chatroom/internal/pkg/utils"
	"io"
	"net"
)

// Processor 先创建一个Processor 的结构体体
type Processor struct {
	Conn net.Conn
}

//编写函数, 根据客户端发送消息的种类不同，决定调用哪个函数来处理
func (process *Processor) processMes(mes *message.Message) (err error) {

	fmt.Println("mes=", mes)
	up := &UserProcess{
		Conn: process.Conn,
	}

	switch mes.Type {
	case message.LoginMesType:
		//处理登录登录
		err = up.ServerProcessLogin(mes)

	case message.RegisterMesType:
		//处理注册
		err = up.ServerProcessRegister(mes) // type : data

	case message.SmsMesType:
		//创建一个 SmsProcess 实例完成转发群聊消息.
		smsProcess := &SmsProcess{}
		smsProcess.SendGroupMes(mes)

	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (process *Processor) ServerProcess() (err error) {

	//循环获取客户端发送的信息
	for {
		//创建一个Transfer 实例完成读包任务
		tf := &utils.Transfer{
			Conn: process.Conn,
		}

		//这里将读取数据包的操作，直接封装成一个函数 readPkg(), 返回 Message, Err
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出..")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}
		err = process.processMes(&mes)
		if err != nil {
			return err
		}
	}
}
