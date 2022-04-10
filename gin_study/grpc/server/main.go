package main

import (
	"context"
	"fmt"
	hello_grpc "gin_study/grpc/pb"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{
		Message: "我是从服务端返回的grpc内容",
	}, nil
}

func main() {
	listen, _ := net.Listen("tcp", ":8888")
	newServer := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(newServer, &server{})
	newServer.Serve(listen)
}
