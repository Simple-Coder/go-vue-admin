package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8888")
	listener, _ := net.ListenTCP("tcp", addr)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			// handle error
			fmt.Println(err)
			return
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn *net.TCPConn) {
	fmt.Print(conn.RemoteAddr().String() + "连接了一下:")
	bytes := make([]byte, 1024)
	read, _ := conn.Read(bytes)
	fmt.Println(string(bytes[0:read]))
}
