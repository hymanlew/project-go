package process

import (
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
