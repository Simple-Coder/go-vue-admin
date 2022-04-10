package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct {
}
type Req struct {
	NumOne int
	NumTwo int
}
type Res struct {
	Num int
}

//第二个参数必须是指针
func (s *Server) Add(req Req, res *Res) error {
	time.Sleep(1 * time.Second)
	res.Num = req.NumOne + req.NumTwo
	return nil
}

func main() {
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("错误")
	}
	http.Serve(listen, nil)
}
