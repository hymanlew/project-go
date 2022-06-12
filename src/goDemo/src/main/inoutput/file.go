package inoutput

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

/*
文件在程序中是以流的形式来操作的。
os.File 封装了所有文件相关的操作，File 是一个结构体。
- os 包提供了操作系统函数的不依赖平台的接口。设计为 Unix 风格，虽然错误处理是 go 风格。失败调用会返回错误值而非错误码。
通常错误值里包含更多信息。例如，如果使用一个文件名的调用（如 Open、Stat）失败了，打印错误时会包含该文件名，错误类型将为 *PathError，
其内部可以解包获得更多信息。
*/
func createClose() {
	// file 叫 file对象、file指针、或者 file 文件句柄
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//输出文件，看出 file 就是一个指针 *File
	fmt.Printf("file=%v", file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}

func read() {
	//1，带缓冲区的读取文件方式：使用 os.Open、file.Close、bufio.NewReader()、reader.ReadString 函数和方法。
	file, err := os.Open("d:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//当函数退出时，要及时的关闭file，否则会有内存泄漏
	defer file.Close()

	// 创建一个 *Reader 指针，是带缓冲的，且默认的缓冲区为 4096
	/*
		const (
			defaultBufSize = 4096
		)
	*/
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF {                  // io.EOF表示文件的末尾
			break
		}
		//因为文件内容中已经含有了换行符，所以不需要再使用 println 方法换行了，否则就会出现两个换行
		fmt.Printf(str)
	}
	fmt.Println("文件读取结束...")

	//2，使用 ioutil.ReadFile 一次将整个文件读入到内存中，适用于文件不大的情况。封装了 opent、close 操作函数。
	filePath := "d:/test.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	fmt.Printf("%v", string(content)) // []byte
}

func write() {
	//OpenFile 文件打开函数（一般会使用 Open 或 Create 代替本函数），它使用指定的打开方式（如O_RDONLY等，也可以组合使用）、
	//指定的权限模式（如0666等）打开指定名称的文件。如果操作成功，返回的文件对象可用于 I/O。如果出错，错误底层类型是*PathError。
	//FileMode 代表文件的模式和权限位，该参数用于 linux、unix 系统，在 windows 系统中无效。
	/*
		func OpenFile(name string, flag int, perm FileMode) (file *File, err error)

		const (
		    O_RDONLY int = syscall.O_RDONLY   只读模式打开文件
		    O_WRONLY int = syscall.O_WRONLY   只写模式打开文件
		    O_RDWR   int = syscall.O_RDWR     读写模式打开文件
		    O_APPEND int = syscall.O_APPEND   写操作时将数据附加到文件尾部
		    O_CREATE int = syscall.O_CREAT    如果不存在将创建一个新文件
		    O_EXCL   int = syscall.O_EXCL     和O_CREATE配合使用，文件必须不存在
		    O_SYNC   int = syscall.O_SYNC     打开文件用于同步I/O
		    O_TRUNC  int = syscall.O_TRUNC    如果可能，打开时清空文件
		)

		const (
		    // 单字符是被String方法用于格式化的属性缩写。
		    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
		    ModeAppend                                     // a: 只能写入，且只能写入到末尾
		    ModeExclusive                                  // l: 用于执行
		    ModeTemporary                                  // T: 临时文件（非备份文件）
		    ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
		    ModeDevice                                     // D: 设备
		    ModeNamedPipe                                  // p: 命名管道（FIFO）
		    ModeSocket                                     // S: Unix域socket
		    ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
		    ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
		    ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
		    ModeSticky                                     // t: 只有root/创建者能删除/移动文件

		    // 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
		    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
		    ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
		)

		//这些被定义的位是FileMode最重要的位。另外9个不重要的位为标准 Unix rwxrwxrwx权限（任何人都可读、写、运行）。
		//这些（重要）位的值应被视为公共API的一部分，可能会用于线路协议或硬盘标识：它们不能被修改，但可以添加新的位。
	*/

	filePath := "d:/abc.txt"
	//file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, fs.ModeAppend) //覆盖写入
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, fs.ModeAppend)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()

	//写入时，使用带缓存的 *Writer
	str := "hello,Gardon\r\n" // \r\n 表示换行
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存，因此在调用 WriterString 方法时，其实内容是先写入到缓存，所以需要调用 Flush 方法，将缓冲的数据
	//真正写入到文件中，否则文件中会没有数据!!!
	writer.Flush()
}
