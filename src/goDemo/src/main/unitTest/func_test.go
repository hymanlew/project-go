package unitTest

import (
	"goDemo/src/main/dataType"
	"goDemo/src/main/funct"
	"goDemo/src/main/utils"
	"testing"
)

/*
Go 语言中自带有一个轻量级的测试框架 testing 和自带的 go test 命令来实现单元测试和性能测试。
testing 框架和其他语言中的测试框架类似，可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例。

- 测试用例文件名必须以 _test.go 结尾。
- 测试用例函数必须以 Test 开头，一般来说就是 Test + 被测试的函数名（首字母大写）。
- 比如 TestAddUpper(t *tesing.T) 的形参类型必须是 *testing.T。

- PASS 表示测试用例运行成功，FAIL 表示测试用例运行失败。

//测试单个文件，一定要带上被测试的原文件，不指定则默认是执行所有的测试文件
go test -v  cal_test.go  cal.go
//测试单个方法
go test -v -test.run TestDate

//获取 test 指令帮助信息
go help test
*/

func TestPublic(t *testing.T) {
	// 如果变量名、函数名、常量名首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用。
	// 可简单理解成，首字母大写是公开的，小写是私有的，在 golang 中没有 public，private 等关键字。
	dataType.Data()
}

func TestDate(t *testing.T) {
	err := funct.Monthwork("a")
	if err != nil {
		//t.Fatalf("执行错误，%v\n", err)
		panic(any(err))
	}
	t.Logf("AddUpper(10) 执行正确...")
}

func TestFunc(t *testing.T) {
	i := utils.Lengh("hello 你好")
	if i <= 0 {
		t.Fatalf("error %v", i)
	}
	t.Logf("success %v", i)
}
