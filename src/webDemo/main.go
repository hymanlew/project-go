package main

import (
	"fmt"
	"net/http"
	"time"
)

//创建处理器函数, 参数顺序是固定的，不能调换
func handler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintln(writer, "hello world", request.URL.Path)
	if err != nil {
		return
	}
}

//浏览器测试路径，http://localhost:8080/test
func test1() {
	//调用处理器
	//HandleFunc 是一个适配器，可以将一个普通函数转换为 HTTP 处理器注册并使用。该适配器通过指定并调用了 HTTP 处理器，从而实现了 Handler
	//接口中的 ServeHTTP 方法
	http.HandleFunc("/test", handler)

	//路由
	//ListenAndServe 会监听 TCP addr，并使用 handler 参数调用 server 函数处理接收到的连接。handler 参数一般设为 nil，表示使用 go
	//默认的多路复用处理器 DefaultServeMux
	//多路复用处理器接收到请求后，要根据请求的 URL 判断使用哪个处理器来处理请求，找到后会重定向到对应处理器处理请求
	//如果指向空的 addr，则会默认使用 80 端口进行网络连接
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}

type MyHandler struct{}

func (mh *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "执行自定义的处理器处理请求")
}

//浏览器测试路径，http://localhost:8080/test2
func test2() {
	//自定义请求处理器
	myhandler := MyHandler{}
	http.Handle("/test2", &myhandler)
	http.ListenAndServe(":8080", nil)
}

//浏览器测试路径，http://localhost:8080/test2，因为在 test2 中已经为路径和处理器进行了绑定
func test3() {
	//自定义 HTTP server 并指定请求处理器，使用的则是指定的请求处理器。多个路径则需要定义多个处理器，此时没有多路复用器了
	myhandler := MyHandler{}
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myhandler,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
}

//浏览器测试路径，http://localhost:8080/test2，因为在 test2 中已经为路径和处理器进行了绑定
func test4() {
	//创建自定义多路复用器
	mux := http.NewServeMux()
	mux.HandleFunc("test4", handler)
	http.ListenAndServe(":8080", mux)
}

func main() {
	//test1()
	//test2()
	test3()
}
