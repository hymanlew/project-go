package inoutput

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

func contentCopy() {
	file1Path := "d:/abc.txt"
	file2Path := "e:/kkk.txt"

	//封装了 opent、close 操作函数，没有使用缓冲区
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, fs.ModeAppend)
	if err != nil {
		fmt.Printf("write file error=%v\n", err)
	}
}

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

func fileCopy() {
	//将 d:/flower.jpg 文件拷贝到 e:/abc.jpg
	srcFile := "d:/flower.jpg"
	dstFile := "e:/abc.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Printf("拷贝完成\n")
	} else {
		fmt.Printf("拷贝错误 err=%v\n", err)
	}
}

// CharCount 定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount    int // 记录英文个数
	NumCount   int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	OtherCount int // 记录其它字符的个数
}

func test() {
	//思路: 打开一个文件, 创一个Reader
	//每读取一行，就去统计该行有多少个 英文、数字、空格和其他字符
	//然后将结果保存到一个结构体
	fileName := "e:/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()

	//定义个 CharCount 实例
	var count CharCount
	//创建一个Reader
	reader := bufio.NewReader(file)

	//开始循环的读取fileName的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //读到文件末尾就退出
			break
		}
		//遍历 str，进行统计
		cstr := []rune(str)
		for _, v := range cstr {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	//输出统计的结果看看是否正确
	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
