package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_learn/pb/user"
	"net"
	"time"
)

//取出server
type userServer struct {
	user.UnimplementedWorkServiceServer
}

//挂载
func (*userServer) Work(ctx context.Context, req *user.UserReq) (*user.UserRes, error) {
	name := req.GetName()
	age := req.GetAge()
	res := &user.UserRes{
		Name: "我收到了:" + name,
		Age:  10 + age,
	}
	return res, nil
}
func (*userServer) WorkStreamIn(server user.WorkService_WorkStreamInServer) error {
	for true {
		recv, err := server.Recv()
		fmt.Println(recv)
		if err != nil {
			server.SendAndClose(&user.UserRes{
				Name: "完成了",
				Age:  0,
			})
			break
		}
	}
	return nil
}
func (*userServer) WorkStreamOut(req *user.UserReq, server user.WorkService_WorkStreamOutServer) error {
	name := req.GetName()
	for i := 0; i < 10; i++ {
		if i > 10 {
			break
		}
		time.Sleep(1 * time.Second)
		server.Send(&user.UserRes{Name: "我拿到了！" + name})
	}
	return nil
}
func (*userServer) WorkStreamAll(server user.WorkService_WorkStreamAllServer) error {
	str := make(chan string)
	i := 0
	go func() {
		for true {
			i++
			req, _ := server.Recv()
			fmt.Println("server收到:" + req.GetName())
			if i > 10 {
				str <- "结束了"
				break
			}
			str <- req.Name
		}
	}()
	for true {
		if <-str == "结束了" {
			break
		}
		server.Send(&user.UserRes{Name: <-str})
	}
	return nil
}
func main() {
	l, _ := net.Listen("tcp", ":8888")
	server := grpc.NewServer()
	user.RegisterWorkServiceServer(server, &userServer{})
	server.Serve(l)
}
