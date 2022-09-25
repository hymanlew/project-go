package rpc

import (
	"fmt"
	goproto "google.golang.org/protobuf/proto"
	"webDemo/target/hyman.com/goDemo/rpc/proto"
)

/**
在 GO 中实现 RPC 非常简单，有封装好的官方库和一些第三方库。Go RPC 可以利用 tcp 或 http 来传递数据，可以对传递的数据使用多种类型的
编码方式。

- 其中 go 官方的 net/rpc 库使用 encoding/gob 进行编解码，支持 tcp 或 http 传递数据，但由于其他语言不支持 gob 编解码方式，所以使
用 net/rpc 库实现的 RPC 方法没法进行跨语言调用。

- go 官方还提供了 net/rpc/jsonrpc 的实现，它采用了 json 进行编解码，因而支持跨语言调用。但目前 jsonrpc 库是基于 tcp 协议实现的，
暂不支持使用 http 进行数据传输。

- 另外还有第三方库为 go 提供 RPC 支持，大部分三方 RPC 库都是使用 protobuf 实现并进行数据编解码的，根据 protobuf 声明文件自动生成
RPC 方法与服务注册代码，在 go 中可以很方便的进行 RPC 服务调用。

在 linux 系统中进行模拟测试时，可使用 nc 命令模拟服务端来接收消息。该命令的作用是：
- 实现任意 TCP/UDP 端口的侦听，故可作为 server 以 TCP/UDP 方式监听指定的端口
- 端口的扫描，可作为 client 发起 TCP/UDP 连接
- 机器之间传输文件，网络测速
*/

/**
Protobuf 是 Protocol Buffers 的简称，是一种数据描述语言，是一种轻便高效的结构化数据存储格式，可用于结构化数据串行化，序列化，支持
跨语言传输。
- 对比于 XML、JSON 更小更快，更简单，很适合做数据存储或 RPC 数据交换。
- 支持跨平台、跨语言。
- 消息格式升级和兼容性好。
- 序列化、反序列化速度很快，且快于 json 的处理速度。

使用前需要先安装，https://github.com/protocolbuffers/protobuf/release/tag/xxx，下载对应版本的安装包，之后解压并将解压后的 bin
目录路径添加到统的 path 系统变量中即可。
查看是否安装成功及版本，protoc --version。

在 go 语言中使用 protobuf 工具，需要安装对应插件：
安装命令: go install github.com/golang/protobuf/protoc-gen-go@latest
检查是否安装成功，无输出则代表安装成功: protoc-gen-go

如果提示非内部命令，则需要看看 gopath 目录（go env 命令查看）下 bin 目录下是否有 protoc-gen-go.exe 文件。没有则代表下载失败，需
要重新下载。有则代表下载成功，然后把当前 bin 目录添加到统的 path 系统变量中即可。
*/

/**
protoc --proto_path=IMPORT_PATH --go_out=DST_DIR path/to/file.proto
protoc 命令默认会导入同级目录下的 proto 文件，且 out 路径只能指定为相对路径。如果有多个不同目录，则需要指定 proto_path。

protoc --proto_path=./rpc/proto --go_out=./target ./rpc/proto/*.proto
同一目录包下直接引用，可不指定 proto_path，前提是注意命令当前路径

protoc --go_out=./webDemo/target ./webDemo/rpc/proto/*.proto
protoc -I ../proto-imports -I ./ --go_out=plugins=grpc:./target ./rpc/grpc/*.proto
不同包文件之间相互引用，要注意 import 中的路径，及命令执行的当前路径

*/

func TestProto() {
	fmt.Println("测试 protoc 生成的文件")

	sys := proto.SystemPermission{
		SystemId:     1,
		PermissionId: 2,
		ActionId:     "action",
		//Ids: make([]int32, 10),
		Ids:  []int32{1, 2, 3},
		View: proto.SystemView_SYSTEM_VIEW_DEFAULT,
	}

	fmt.Println(sys.GetSystemId())
	data, _ := goproto.Marshal(&sys)
	fmt.Println(data)

	sys2 := proto.SystemPermission{}
	goproto.Unmarshal(data, &sys2)
	fmt.Printf("%#v\n", sys2)
	fmt.Println(sys2.GetActionId())
}
