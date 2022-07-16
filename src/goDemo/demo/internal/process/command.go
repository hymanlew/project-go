package process

import (
	"flag"
	"fmt"
	"os"
)

func commandLine() {
	fmt.Println("命令行的参数有", len(os.Args))

	//遍历 os.Args 切片，就可以得到所有的命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}

	//使用示例
	//编译构建时指定名字和目录
	//go build -o bin/my.exe  go_code/project/main
	//执行该 exe 文件，然后在后面加上命令行参数即可
}

func flagCommand() {
	//上面的方式是比较原生的方式，对解析参数不是很方便，特别是带有指定参数形式的命令行。
	//比如：cmd>main.exe -f c:/aaa.txt -p 200 -u root 这样形式的命令行，
	//因此 go 提供了 flag 包，可以方便的解析命令行参数，而且参数顺序可以随意。

	//定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接收用户命令行中输入的 -u 后面的参数值
	//"u" 就是指定的 -u 参数
	//"" 默认值
	//"用户名,默认为空" 为参数的说明
	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码,默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名,默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	//这里有一个非常重要的操作转换，必须调用该方法
	flag.Parse()

	//输出结果
	fmt.Printf("user=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}
