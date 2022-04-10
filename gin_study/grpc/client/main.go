package main

import (
	"context"
	"fmt"
	hello_grpc "gin_study/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println(err)
	defer conn.Close()
	client := hello_grpc.NewHelloGRPCClient(conn)
	req, _ := client.SayHi(context.Background(), &hello_grpc.Req{
		Message: "我是客户端发来消息了",
	})
	fmt.Println(req.GetMessage())
}
