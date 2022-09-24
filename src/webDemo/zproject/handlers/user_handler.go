package handlers

import (
	"encoding/json"
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
	//usr, err := dbsql.GetByName(name)
	//if usr != nil {
	//	msg = "登录成功" + name
	//} else {
	//	msg = "登录失败，" + err.Error()
	//}
	msg = "登录成功" + name

	//添加 session
	uuid := utils.CreatUUID()
	session := &model.Session{
		SessionId: uuid,
		User: &dbsql.User{
			Id:   1,
			Name: name,
			Age:  10,
		},
	}
	utils.SessionMaap[session.SessionId] = session

	//创建一个 cookie 用于携带 session 信息
	cookie := http.Cookie{
		Name:     "user",
		Value:    uuid,
		HttpOnly: true,
	}
	http.SetCookie(writer, &cookie)

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

func Shops(writer http.ResponseWriter, request *http.Request) {
	//获取用户登录信息
	uuid := ""
	cookie, _ := request.Cookie("user")
	if cookie != nil {
		uuid = cookie.Value
	}

	msg := ""
	if uuid != "" {
		session := utils.SessionMaap[uuid]
		if session != nil {
			bytes, _ := json.Marshal(session.User)
			msg = string(bytes)
		} else {
			msg = "用户不存在！"
		}
	} else {
		msg = "您未登录！"
	}

	fmt.Fprintln(writer, msg)
}

func Logout(writer http.ResponseWriter, request *http.Request) {
	//获取用户登录信息
	cookie, _ := request.Cookie("user")
	if cookie != nil {
		utils.SessionMaap[cookie.Value] = nil
		cookie.MaxAge = -1
	}

	http.SetCookie(writer, cookie)
	temp := template.Must(template.ParseFiles("zproject/pages/index.html"))
	temp.Execute(writer, "退出成功")
}
