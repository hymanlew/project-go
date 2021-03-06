package main

import (
	"fmt"
	"goDemo/project/chatroom/internal/client/process"
	"os"
)

//定义变量，表示用户id, 用户密码, 用户名称
var userId int
var userPwd string
var userName string

func main() {

	//接收用户的选择
	var key int

	//判断是否还继续显示菜单
	var loop = true

	for loop {
		fmt.Println("----------------欢迎登陆多人聊天系统------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")

			//必须要加 \n 即捕捉到回车的输入
			fmt.Println("请输入用户的id")
			_, err := fmt.Scanf("%d\n", &userId)
			if err != nil {
				fmt.Printf("输入有误，err = %v", err)
				continue
			}
			fmt.Println("请输入用户的密码")
			_, err = fmt.Scanf("%s\n", &userPwd)
			if err != nil {
				fmt.Printf("输入有误，err = %v", err)
				continue
			}

			//完成登录, 创建一个 UserProcess 的实例
			up := &clientProcess.UserProcess{}
			err = up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("登录失败")
				continue
			} else {
				fmt.Println("登录成功")
			}
			loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名字(nickname):")
			fmt.Scanf("%s\n", &userName)

			//调用 UserProcess，完成注册的请求
			up := &clientProcess.UserProcess{}
			err := up.Register(userId, userPwd, userName)
			if err != nil {
				fmt.Println("注册失败")
				continue
			} else {
				fmt.Println("注册成功")
			}
			loop = false
		case 3:
			fmt.Println("退出系统")
			//loop = false
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
}
