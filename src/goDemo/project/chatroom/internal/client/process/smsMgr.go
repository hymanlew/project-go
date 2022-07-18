package clientProcess

import (
	"encoding/json"
	"fmt"
	"goDemo/project/chatroom/internal/pkg/message"
)

func outputGroupMes(mes *message.Message) {
	//显示即可，反序列化 mes.Data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}

	//显示信息
	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()

}
