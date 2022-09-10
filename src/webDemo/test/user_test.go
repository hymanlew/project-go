package test

import (
	"fmt"
	"testing"
	"webDemo/dbsql"
)

//TestMain 函数是在其他所有测试函数之前执行，可用于初始化操作
func TestMain(m *testing.M) {
	fmt.Println("测试初始化操作：")
	//必须调用 run 函数，程序才能自动执行其他测试函数
	m.Run()
}

func TestUserAdd(t *testing.T) {
	t.Log("测试新增用户：")
	usr := &dbsql.User{
		1,
		"hyman",
		"123456",
		18,
	}
	usr.Add()
	usr.Add2()
}
