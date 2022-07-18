package clientProcess

import (
	"fmt"
	"goDemo/project/chatroom/internal/pkg/message"
	"goDemo/project/chatroom/internal/pkg/model"
)

//客户端要维护的map
var onlineUsers map[int]*model.User = make(map[int]*model.User, 10)

// CurUser 我们在用户登录成功后，完成了对 CurUser 初始化
var CurUser model.CurUser

//在客户端显示当前在线的用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUsers {
		//如果不显示自己.
		fmt.Println("用户id:\t", id)
	}
}

//编写一个方法，处理返回的 NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	//适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &model.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
