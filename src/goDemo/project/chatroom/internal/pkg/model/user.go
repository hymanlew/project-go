package model

// User 定义一个用户的结构体
type User struct {

	//确定字段信息
	//为了序列化和反序列化成功，必须保证用户信息的 json 字符串的 key 和结构体字段对应的 tag 名字一致
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //用户状态
	Sex        string `json:"sex"`        //性别
}
