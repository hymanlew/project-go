package utils

import (
	"bufio"
	"errors"
	"io"
	"os"
)

func FileCopy(srcPath string, desPath string) (size int64, err error) {

	var srcFile *os.File
	if _, err := os.Stat(srcPath); err != nil || os.IsNotExist(err) {
		return 0, errors.New("源文件不存在")
	} else {
		srcFile, err = os.Open(srcPath)
		if err != nil {
			return 0, errors.New("源文件读取失败")
		}
	}

	desFile, err := os.OpenFile(desPath, os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return 0, errors.New("目标文件打开或创建失败")
	}

	defer srcFile.Close()
	defer desFile.Close()

	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(desFile)
	return io.Copy(writer, reader)
}
