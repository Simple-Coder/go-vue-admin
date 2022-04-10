package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Req struct {
	NumOne int
	NumTwo int
}
type Res struct {
	Num int
}

func main() {
	req := Req{
		NumOne: 1,
		NumTwo: 2,
	}
	client, err := rpc.DialHTTP("tcp", "localhost:8888")
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}

	var res Res
	//同步
	//client.Call("Server.Add", req, &res)
	call := client.Go("Server.Add", req, &res, nil)
	for true {
		select {
		case <-call.Done:
			fmt.Println(res)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("正在等待")
		}
	}
	fmt.Println(res)
}
