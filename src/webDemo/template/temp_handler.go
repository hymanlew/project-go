package http

import (
	"fmt"
	"html/template"
	"net/http"
)

//使用 GO 的 web 模板引擎，需要以下两个步骤：
//- 对文本格式的模板源进行语法分析，创建一个经过语法分析的模板结构，其中模板源即可以是一个字符串，也可以是模板文件中包含的内容。
//- 执行经过语法分析的模板，将 responseWriter 和模板所需的动态数据传递给模板引擎，被调用的模板引擎会把模板和传入的数据结合起来，生
//成最终的 HTML，并将该 HTML 传递给 responseWriter。

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "执行 HTML 模板请求", request.URL.Path)

	//解析模板文件
	temp, _ := template.ParseFiles("template.html")
	//执行模板，动态数据替换。且该方法只能作用于一个模板的情况下，如果有多个模板，则只会作用到第一个模板中
	temp.Execute(writer, "hello template")

	//该方法作用同上，但它是作用于指定 name 的模板产生的输出
	temp, _ = template.ParseFiles("template.html", "temp.html")
	temp.ExecuteTemplate(writer, "temp.html", "hello template2")
}

func test() {
	http.HandleFunc("/temp", handler)
	http.ListenAndServe(":8080", nil)
}
