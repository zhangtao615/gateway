package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			break
		}
		str := string(buf[:n])
		fmt.Printf("receive from client, data: %v\n", str)
	}
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", "0:0:0:0:8081")

	if err != nil {
		fmt.Printf("listen failed, err: %v\n", err)
		return
	}
	
	// 建立套接字层
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err: %v\n", err)
			continue
		}

		// 创建处理协程
		go process(conn)
	}
}