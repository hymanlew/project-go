package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webDemo/dbsql"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "执行 POST 请求", request.URL.Path)
	fmt.Fprintf(writer, "请求体中，长度 %d", request.ContentLength)

	body := make([]byte, 1024)
	request.Body.Read(body)
	fmt.Fprintf(writer, "请求体中，内容为: %v", string(body))

	//Form 字段只有在调用 request 的 parseForm 方法后才有效，才会被赋值。因为在客户端，会忽略请求中的本字段而使用 body 替代。
	//parseForm 会解析 URL 中的查询字符串，并将解析结果更新到 Form 字段。
	//对于 POST/PUT 请求，parseForm 还会将 body 作为表单解析，并将结果更新到 postForm、Form 字段。解析结果过程中，POST/PUT 的
	//请求主体优先于 URL 查询字符串。
	//如果请求体的大小没有被 MaxBytesReader 函数限制，其大小默认限制为 10 MB。

	//如果 form 表单 URL 地址中也有与 form 表单参数名相同的请求参数，则这两种参数值 Form 方法都可以得到。并且表单中的请求参数值是
	//在 URL 参数值的前面
	request.ParseForm()
	fmt.Fprintf(writer, "请求体中，内容为: %v", request.Form)
	fmt.Fprintf(writer, "POST 请求体中，内容为: %v", request.PostForm)

	//ParseMultipartForm 方法会自动调用 parseForm 方法，因此不需要重复调用 parseForm 方法。
	request.ParseMultipartForm(20480)
	fmt.Fprintf(writer, "POST 请求体中，文件类型内容为: %v", request.MultipartForm)

	//该方法用于从表单中获取某一个参数的值，POST/PUT 请求体中的同名参数优先 URL 地址中的参数。且此函数会自动调用 ParseMultipartForm
	//和 parseForm 方法，因此不需要重复调用 parseForm 方法。
	str := request.FormValue("abc")
	str = request.PostFormValue("abc")
	fmt.Fprintf(writer, str)
}

func postHandler(writer http.ResponseWriter, request *http.Request) {
	//浏览器会自动解析 Content-Type: text/plain
	writer.Write([]byte("返回字符串"))

	writer.Header().Set("Content-Type", "application/json")
	user := dbsql.User{
		1,
		"hyman",
		"",
		28,
	}
	usrJson, _ := json.Marshal(user)
	writer.Write(usrJson)

	//302 重定向操作
	writer.Header().Set("Location", "www.baidu.com")
	writer.WriteHeader(302)
}

func cookieHandler(writer http.ResponseWriter, request *http.Request) {
	//HttpOnly:
	//控制 Cookie 的内容是否可以被 JavaScript 访问到。设置为 true 可以防止 XSS 攻击。默认 HttpOnly 为false, 表示客户端可以通过js获取。

	//Path:
	//用于设置 Cookie 的访问范围，默认为 ”/” 表示当前项目下所有都可以访问。也可以对 Path 设置路径，表示此路径及子路径内容都可以访问。

	//Expires:
	//Cookie 默认的存活时间是浏览器开启时间, 当浏览器关闭后, Cookie 才失效。目前 chrome 等主流浏览器都使用 MaxAge 设置 Cookie 的有效
	//时间。但像 IE6,7,8 和其他浏览器不支持 MaxAge, 所以还是使用 Expires。
	cookie := http.Cookie{
		Name:     "cookieName",
		Value:    "value",
		HttpOnly: true,
	}

	//服务器返回 cookie 给客户端，之后每次同一客户端访问时，其浏览器会自动在请求头中发送 cookie 消息
	writer.Header().Set("cookie", cookie.String())
	writer.Header().Add("cookie", cookie.String()+"_2")
	http.SetCookie(writer, &cookie)
}

func cookieGetHandler(writer http.ResponseWriter, request *http.Request) {
	cookie := request.Header.Get("Cookie")
	fmt.Printf("获取到的 cookies: %v", cookie)

	cookie2, _ := request.Cookie("cookieName")
	fmt.Printf("获取到的 cookies: %v", cookie2)
}

func HttpTest() {
	http.HandleFunc("/post", handler)
	http.HandleFunc("/post2", postHandler)
	http.ListenAndServe(":8080", nil)
}
