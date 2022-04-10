package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//fmt.Println(err)
	defer conn.Close()

	//client := user.NewWorkServiceClient(conn)

	//传统
	//work, err := client.Work(context.Background(), &user.UserReq{
	//	Name: "张三...",
	//	Age:  18,
	//})

	//客户端stream
	//c2, _ := client.WorkStreamIn(context.Background())
	//for i := 0; i < 10; i++ {
	//	if i > 10 {
	//		recv, _ := c2.CloseAndRecv()
	//		fmt.Println(recv)
	//		break
	//	}
	//	time.Sleep(1 * time.Second)
	//	c2.Send(&user.UserReq{
	//		Name: "我是进来的信息",
	//		Age:  10,
	//	})
	//}

	// 响应 stream
	//c3, _ := client.WorkStreamOut(context.Background(), &user.UserReq{Name: "我是张三"})
	//for true {
	//	recv, err := c3.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(recv)
	//}

	//c4, _ := client.WorkStreamAll(context.Background())
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//go func() {
	//	for {
	//		//time.Sleep(1 * time.Second)
	//		err := c4.Send(&user.UserReq{Name: "我是张三"})
	//		if err != nil {
	//			fmt.Println(err)
	//			wg.Done()
	//			break
	//		}
	//	}
	//}()
	//
	//go func() {
	//	for true {
	//		recv, err := c4.Recv()
	//		if err != nil {
	//			fmt.Println(err)
	//			wg.Done()
	//			break
	//		}
	//		fmt.Println("client收到：" + recv.GetName())
	//	}
	//}()
	//wg.Wait()
}
