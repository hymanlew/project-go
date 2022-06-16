package inoutput

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type A struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float32 `json:"score"`
}

func Store(a *A) {
	fmt.Printf("准备写入的数据 == %v\n", *a)
	bytes, err := json.Marshal(a)
	if err != nil {
		panic(any("序列化失败"))
	}
	file, err := os.Create("E:\\project-go\\src\\test.txt")
	if err != nil {
		panic(any("创建文件失败"))
	}
	writer := bufio.NewWriter(file)
	i, err := writer.Write(bytes)
	if i == 0 {
		panic(any("写入文件失败"))
	}
	writer.Flush()
	file.Close()
}

func ReStore(a *A) {
	file, err := os.Open("E:\\project-go\\src\\test.txt")
	if err != nil {
		panic(any("打开文件失败"))
	}
	result := make([]byte, 0, 1024)
	value := make([]byte, 1024, 1024)
	reader := bufio.NewReader(file)
	for {
		data, err := reader.Read(value)
		if data == 0 || err == io.EOF {
			fmt.Println("== 读取内容完毕 ==")
			break
		}
		result = append(result, value[:data]...)
	}
	file.Close()
	json.Unmarshal(result, a)
	fmt.Printf("读取到的数据 == %v\n", *a)
}
