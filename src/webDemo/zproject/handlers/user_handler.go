package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"webDemo/dbsql"
	"webDemo/zproject/model"
	"webDemo/zproject/utils"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	//获取用户登录信息
	name := request.PostFormValue("name")
	pass := request.PostFormValue("pass")
	fmt.Printf("name = %v, pass =  %v", name, pass)

	var msg string
	usr, err := dbsql.GetByName(name)
	if usr != nil {
		msg = "登录成功" + name
	} else {
		msg = "登录失败，" + err.Error()
	}

	//添加 session
	session := &model.Session{
		SessionId: ",",
		User: &dbsql.User{
			Id:   1,
			Name: name,
			Age:  10,
		},
	}
	utils.SessionMaap[session.SessionId] = session

	temp := template.Must(template.ParseFiles("zproject/pages/view/msg.html"))
	temp.Execute(writer, msg)
}

func Regist(writer http.ResponseWriter, request *http.Request) {
	//获取用户登录信息
	name := request.PostFormValue("name")
	pass := request.PostFormValue("pass")
	fmt.Printf("name = %v, pass =  %v", name, pass)

	temp := template.Must(template.ParseFiles("zproject/pages/view/msg.html"))
	temp.Execute(writer, "注册成功: "+name)
}
