package serverProcess

import (
	"encoding/json"
	"fmt"
	"goDemo/project/chatroom/internal/pkg/message"
	"goDemo/project/chatroom/internal/pkg/model"
	"goDemo/project/chatroom/internal/pkg/service"
	"goDemo/project/chatroom/internal/pkg/utils"
	"net"
)

type UserProcess struct {
	//用户保持的连接
	Conn net.Conn
	//表示该 Conn 是哪个用户
	UserId int
}

// NotifyOthersOnlineUser 通知其他所有在线的用户，userId 我上线了
func (userProcess *UserProcess) NotifyOthersOnlineUser(userId int) {

	//遍历 onlineUsers, 然后一个一个的发送 NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤到自己
		if id == userId {
			continue
		}
		//开始通知【单独的写一个方法】
		up.NotifyMeOnline(userId)
	}
}

func (userProcess *UserProcess) NotifyMeOnline(userId int) {

	//组装 NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给 mes.Data
	mes.Data = string(data)

	//对mes再次序列化，准备发送.
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送,创建我们Transfer实例，发送
	tf := &utils.Transfer{
		Conn: userProcess.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
}

func (userProcess *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	//从 mes 中取出 mes.Data，并直接反序列化成 RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	//到 redis 数据库完成注册
	err = service.NewUserDao(utils.GetPool()).Register(&registerMes.User)
	if err != nil {
		if err == model.ErrorUserExists {
			registerResMes.Code = 505
			registerResMes.Error = model.ErrorUserExists.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//将 data 赋值给 resMes，并对 resMes 进行序列化，准备发送
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//发送 data, 我们将其封装到 writePkg 函数
	tf := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// ServerProcessLogin 编写函数，专门处理登录请求
func (userProcess *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

	//先从 mes 中取出 mes.Data，并直接反序列化成 LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//声明一个 LoginResMes，并完成赋值
	var loginResMes message.LoginResMes

	//并到 redis 数据库去完成验证
	userDao := service.NewUserDao(utils.GetPool())
	user, err := userDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ErrorUserNotexists {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ErrorUserPwd {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {

		//用户登录成功，就放入到 userMgr 中
		loginResMes.Code = 200
		userProcess.UserId = loginMes.UserId
		userMgr.AddOnlineUser(userProcess)

		//通知其它的在线用户，我上线了
		userProcess.NotifyOthersOnlineUser(loginMes.UserId)

		//将当前在线用户的 id 放入到 loginResMes.UsersId
		//遍历 userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user, "登录成功")
	}

	//将 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//声明一个 resMes, 并将 data 赋值给 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	resMes.Data = string(data)

	//对 resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//发送 data, 我们将其封装到 writePkg 函数
	tf := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = tf.WritePkg(data)
	return
}
