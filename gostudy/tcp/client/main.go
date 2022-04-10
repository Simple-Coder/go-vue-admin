package main

import (
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8888")
	conn, _ := net.DialTCP("tcp", nil, addr)
	//手写输入
	//os.Stdin
	defer conn.Close()
	conn.Write([]byte("我申请出站"))

}
