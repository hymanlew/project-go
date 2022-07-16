package utils

import (
	"fmt"
	"strconv"
	"strings"
)

var stre = "hello ya"
var strc = "hello 你好呀"
var strn = "10"

func util() {
	//长度，返回字节数
	slen := len(stre)
	fmt.Println(slen)

	//字符串遍历，同时处理有中文的问题
	r := []rune(strc)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	//字符串转整数 a -- int
	i := ParseToInt(strn)
	fmt.Println(i)

	//整数转字符串 int -- a
	s := strconv.Itoa(123)
	fmt.Println(s)

	//字符串转 []byte，输出为 ASCII 编码
	bytes := []byte(stre)
	fmt.Println(bytes)

	//[]byte 转字符串，将 ASCII 编码转字符
	s = string([]byte{})
	fmt.Println(s)

	//10进制转 2, 8, 16进制
	s = strconv.FormatInt(10, 2)
	fmt.Println(s)

	//查找子串是否在指定的字符串中
	t := strings.Contains(stre, strc)
	fmt.Println(t)

	//统计一个字符串有几个指定的子串
	i = strings.Count(strc, strn)
	fmt.Println(i)

	//不区分大小写的字符串比较
	//== 是只区分字母大小写的，内容完全相同的字符串
	t = strings.EqualFold(strn, strc)
	fmt.Println(t)
	fmt.Println(strc == strn)

	//返回子串在字符串第一次出现的 index 值，没有则返回-1
	i = strings.Index(strn, strc)
	fmt.Println(i)
	i = strings.LastIndex(strc, strn)
	fmt.Println(i)

	//将指定的子串替换成另外一个子串，n 可以指定替换的个数，如果 n=-1 表示全部替换，返回一个新字符串，原先字符串不变
	s = strings.Replace(strn, "a", "b", 1)
	fmt.Println(s)

	//按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
	strArr := strings.Split(strn, ",")
	fmt.Println(strArr)

	strings.ToLower("")
	strings.ToUpper("")

	//将字符串左右两边的空格去掉
	strings.TrimSpace(strn)
	//将字符串左右两边指定的字符去掉，并且可以指定多个字符
	strings.Trim(strn, ",")
	strings.Trim(stre, "!,")
	strings.TrimLeft(strn, ",")
	strings.TrimRight(strn, ",")

	//判断字符串是否以指定的字符串开头
	strings.HasPrefix(stre, "a")
	strings.HasSuffix(strn, "a")
}

func Lengh(str string) int {
	n := []rune(str)
	return len(n)
}

func ParseToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	} else {
		return i
	}
}

func ParseToStr(num int) string {
	s := strconv.Itoa(num)
	return s
}
