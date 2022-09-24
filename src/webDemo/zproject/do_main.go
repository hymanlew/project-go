package zproject

import (
	"net/http"
	"text/template"
	"webDemo/zproject/handlers"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	//解析模板
	//该函数用于包装返回的模板，它会在 err 非 nil 时 panic，一般用于变量初始化
	temp := template.Must(template.ParseFiles("zproject/pages/index.html"))
	temp.Execute(writer, "")
}

func DoMain() {
	//处理静态资源
	//StripPrefix 函数用于将 URL 地址中指定的前缀去除后，再交由处理器。如果没有指定前缀，则会响应 404
	//FileServer 函数返回一个使用 FileSystem 接口 root 提供文件访问服务的 HTTP 处理器。要使用 OS 的 FileSystem 接口实现，可使用 http.Dir
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("zproject/pages/static"))))
	http.Handle("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("zproject/pages/view"))))
	http.HandleFunc("/shop/index/", handler)

	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/regist", handlers.Regist)
	http.HandleFunc("/getShops", handlers.Shops)
	http.HandleFunc("/logout", handlers.Logout)

	http.ListenAndServe(":8080", nil)
}
