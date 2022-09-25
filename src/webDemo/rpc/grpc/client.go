package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"webDemo/target/hyman.com/goDemo/rpc/proto/userService"
)

func TestgRPCClient() {
	//连接到服务器
	//从输入的证书文件中为客户端构造 TLS 凭证
	//grpc.Dial("127.0.0.1:8080", credentials.NewClientTLSFromFile(cerFile,"tls"))
	//若没有证书文件，则使用下面，配置连接级别的安全凭证（TLS/SSL），并连接服务器
	grpcClient, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}

	//注册客户端
	client := userService.NewUserServiceClient(grpcClient)

	res, err := client.ListSystems(context.Background(), &userService.UserRequest{
		Id:   1,
		Name: "abc",
		Age:  28,
	})
	fmt.Printf("%v\n", res)
	fmt.Println(res.GetUsers()[0])
}
