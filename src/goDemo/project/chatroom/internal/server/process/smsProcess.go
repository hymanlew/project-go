package serverProcess

import (
	"encoding/json"
	"fmt"
	"goDemo/project/chatroom/internal/pkg/message"
	"goDemo/project/chatroom/internal/pkg/utils"
	"net"
)

type SmsProcess struct {
	//..[暂时不需字段]
}

// SendGroupMes 写方法转发消息
func (sms *SmsProcess) SendGroupMes(mes *message.Message) {

	//遍历服务器端的 onlineUsers map[int]*UserProcess,
	//将消息转发取出，取出 mes 的内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//还需要过滤到自己, 即不要再发给自己
		if id == smsMes.UserId {
			continue
		}
		sms.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (sms *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {

	//创建一个Transfer 实例，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
	}
}
