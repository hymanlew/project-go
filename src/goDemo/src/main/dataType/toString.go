package dataType

import (
	"fmt"
	"strconv"
)

func tostring()  {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string //空的str

	//转换 string, 使用 fmt.Sprintf 方法
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q\n", str, str)

	//第二种方式 strconv 函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	// strconv.FormatFloat(num4, 'f', 10, 64)
	// 'f' 浮点数格式
	// 10：表示小数位保留10位
	// 64 :表示这个小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)

	//strconv包中有一个函数Itoa
	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str=%q\n", str, str)
}

func strFormat()  {
	//通用：
	//%v	值的默认格式表示
	//%+v	类似%v，但输出结构体时会添加字段名
	//%#v	值的Go语法表示
	//%T	值的类型的Go语法表示
	//%%	百分号

	//布尔值：
	//%t	单词true或false

	//整数：
	//%b	表示为二进制
	//%c	该值对应的unicode码值
	//%d	表示为十进制
	//%o	表示为八进制
	//%q	该值对应的单引号括起来的字符字面值，必要时会采用安全的转义表示。如 string 会带双引号输出
	//%x	表示为十六进制，使用a-f
	//%X	表示为十六进制，使用A-F
	//%U	表示为Unicode格式：U+1234，等价于"U+%04X"

	//浮点数与复数的两个组分：
	//%b	无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
	//%e	科学计数法，如-1234.456e+78
	//%E	科学计数法，如-1234.456E+78
	//%f	有小数部分但无指数部分，如123.456
	//%F	等价于%f
	//%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	//%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）

	//字符串和[]byte：
	//%s	直接输出字符串或者[]byte
	//%q	该值对应的单引号括起来的字符字面值，必要时会采用安全的转义表示。如 string 会带双引号输出
	//%x	每个字节用两字符十六进制数表示（使用a-f）
	//%X	每个字节用两字符十六进制数表示（使用A-F）

	//指针：
	//%p	表示为十六进制，并加上前导的0x

	//没有%u。整数如果是无符号类型自然输出也是无符号的。类似的，也没有必要指定操作数的尺寸（int8，int64）。
}

func stringTo()  {
	var str string = "true"
	var b bool

	// 1. 该函数会返回两个值 (value bool, err error)
	// 2. 因为只想获取到 value bool ,不想获取 err 所以我使用_忽略
	b , _ = strconv.ParseBool(str)
	fmt.Printf("b type %T  b=%v\n", b, b)

	var str2 string = "1234590"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T  n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)

	//注意：
	var str4 string = "hello"
	var n3 int64 = 11
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type %T n3=%v\n", n3, n3)
}
