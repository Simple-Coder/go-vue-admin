package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_learn/pb/user"
	"net"
	"net/http"
	"sync"
)

//取出server
type userServer struct {
	user.UnimplementedWorkServiceServer
}

//挂载
func (*userServer) Work(ctx context.Context, req *user.UserReq) (*user.UserRes, error) {
	name := req.GetName()
	res := &user.UserRes{
		Name: "我收到了" + name,
	}
	return res, nil
}

//func (*userServer) WorkStreamIn(server user.WorkService_WorkStreamInServer) error {
//	for true {
//		recv, err := server.Recv()
//		fmt.Println(recv)
//		if err != nil {
//			server.SendAndClose(&user.UserRes{
//				Name: "完成了",
//				Age:  0,
//			})
//			break
//		}
//	}
//	return nil
//}
//func (*userServer) WorkStreamOut(req *user.UserReq, server user.WorkService_WorkStreamOutServer) error {
//	name := req.GetName()
//	for i := 0; i < 10; i++ {
//		if i > 10 {
//			break
//		}
//		time.Sleep(1 * time.Second)
//		server.Send(&user.UserRes{Name: "我拿到了！" + name})
//	}
//	return nil
//}
//func (*userServer) WorkStreamAll(server user.WorkService_WorkStreamAllServer) error {
//	str := make(chan string)
//	i := 0
//	go func() {
//		for true {
//			i++
//			req, _ := server.Recv()
//			fmt.Println("server收到:" + req.GetName())
//			if i > 10 {
//				str <- "结束了"
//				break
//			}
//			str <- req.Name
//		}
//	}()
//	for true {
//		if <-str == "结束了" {
//			break
//		}
//		server.Send(&user.UserRes{Name: <-str})
//	}
//	return nil
//}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go registerGateway(&wg)
	go registerGrpc(&wg)
	wg.Wait()
}
func registerGrpc(wg *sync.WaitGroup) {
	l, _ := net.Listen("tcp", ":8888")
	server := grpc.NewServer()
	user.RegisterWorkServiceServer(server, &userServer{})
	server.Serve(l)
	wg.Done()
}

func registerGateway(wg *sync.WaitGroup) {
	conn, _ := grpc.DialContext(context.Background(),
		"0.0.0.0:8888",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	mux := runtime.NewServeMux()
	user.RegisterWorkServiceHandler(context.Background(), mux, conn)

	gwServer := &http.Server{
		Handler: mux,
		Addr:    ":8090",
	}
	gwServer.ListenAndServe()

	wg.Done()
}
