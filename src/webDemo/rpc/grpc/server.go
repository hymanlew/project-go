package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"webDemo/target/hyman.com/goDemo/rpc/proto"
	"webDemo/target/hyman.com/goDemo/rpc/proto/userService"
)

// GrpcDemo 自定义实现类实现 userService 的相关接口
type GrpcDemo struct{}

func (demo *GrpcDemo) ListSystems(c context.Context, req *userService.UserRequest) (res *userService.ListUserResponse, err error) {
	fmt.Printf("%v", req)

	user := userService.UserRequest{
		Id:   1,
		Name: "abc",
		Age:  28,
	}
	res = &userService.ListUserResponse{
		Users:     []*userService.UserRequest{&user},
		TotalSize: 1,
	}
	return res, nil
}

func (demo *GrpcDemo) CreateSystem(c context.Context, req *userService.UserRequest) (*userService.ListUserResponse, error) {
	fmt.Printf("%v", req)
	return &userService.ListUserResponse{}, nil
}

func (demo *GrpcDemo) UpdateSystemPermissions(c context.Context, req *proto.UpdatePermissionRequest) (*proto.Permission, error) {
	fmt.Printf("%v", req)
	return &proto.Permission{}, nil
}

func (demo *GrpcDemo) DeleteSystemPermissions(c context.Context, req *proto.DeletePermissionRequest) (*proto.Permission, error) {
	fmt.Printf("%v", req)
	return nil, nil
}

func (demo *GrpcDemo) VerifyPermission(c context.Context, req *proto.UpdatePermissionRequest) (*proto.Permission, error) {
	fmt.Printf("%v", req)
	return &proto.Permission{}, nil
}

func TestgRPCServer() {
	//初始化一个 grpc 对象
	grpcServer := grpc.NewServer()
	//注册服务
	userService.RegisterUserServiceServer(grpcServer, new(GrpcDemo))

	//监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	//启动服务
	grpcServer.Serve(listener)
}
