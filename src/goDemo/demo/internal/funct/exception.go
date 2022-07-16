package funct

import (
	"errors"
	"fmt"
)

func makeErr() {
	//使用 defer + recover 匿名函数来捕获和处理异常
	defer func() {
		//recover()内置函数，可以捕获到异常
		if err := recover(); err != any(nil) {
			fmt.Println("发生错误 ~ ", err)
			fmt.Println("发送补偿等操作~~~")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

//函数去读取以配置文件init.conf的信息
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误..")
	}
}

func Test7() {
	makeErr()
	fmt.Println("接着执行程序中 1 ~~~")

	err := readConf("config2.ini")
	if err != nil {
		//如果读取文件发送错误，就输出这个错误，并终止程序
		panic(any(err))
	}
	fmt.Println("接着执行程序中 2 ~~~")
}
