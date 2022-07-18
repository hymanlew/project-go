package serverProcess

import (
	"fmt"
)

//因为 UserMgr 实例在服务器端有且只有一个，且在很多的地方都会使用到，因此将其定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// AddOnlineUser 完成对 onlineUsers 添加
func (usrMgr *UserMgr) AddOnlineUser(up *UserProcess) {
	usrMgr.onlineUsers[up.UserId] = up
}

// DelOnlineUser 删除
func (usrMgr *UserMgr) DelOnlineUser(userId int) {
	delete(usrMgr.onlineUsers, userId)
}

// GetAllOnlineUser 返回当前所有在线的用户
func (usrMgr *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return usrMgr.onlineUsers
}

// GetOnlineUserById 根据id返回对应的值
func (usrMgr *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {

	//如何从map取出一个值，带检测方式
	up, ok := usrMgr.onlineUsers[userId]
	if !ok { //说明，你要查找的这个用户，当前不在线。
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
