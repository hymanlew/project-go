package clientProcess

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"goDemo/project/chatroom/internal/pkg/message"
	"goDemo/project/chatroom/internal/pkg/model"
	"goDemo/project/chatroom/internal/pkg/utils"
	"net"
	"os"
)

type UserProcess struct {
	//暂时不需要字段..
}

func (usrProcess *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

	//链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//准备通过 conn 发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//创建一个 Transfer 实例, 发送 data 给服务器端
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
	}

	// mes 就是 RegisterResMes
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将 mes 的 Data 部分反序列化成 RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功, 你重新登录一把")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

func (usrProcess *UserProcess) Login(userId int, userPwd string) (err error) {

	//链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//将 loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//把 data 赋给 mes.Dat a字段, 并将 mes 进行序列化化
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//自定义通讯协议，发送数据包
	//先获取到 data 的长度 -> 转成一个表示长度的 byte 切片，使用 4 个字节
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	fmt.Printf("客户端，发送消息的长度=%d 内容=%s", len(data), string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	//休眠20
	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20..")

	//这里还需要处理服务器端返回的消息
	//创建一个Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将 mes的Data部分反序列化成 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		fmt.Println("登录成功")

		//可以显示当前在线用户列表, 遍历 loginResMes.UsersId
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UsersId {
			//如果要求不显示自己在线, 则跳过自己即可
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)

			//完成客户端的 onlineUsers 初始化
			user := &model.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		//这里还需要在客户端启动一个协程，该协程保持和服务器端的通讯
		//如果服务器有数据推送给客户端，则接收并显示在客户端的终端
		go serverProcessMes(conn)

		//显示登录成功的菜单[循环]..
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
