package main

import (
	"fmt"
	"net"
)

func main() {
	coon, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})

	if err != nil {
		fmt.Printf("connect failed, err: %v\n", err)
		return
	}

	for i := 0; i < 100; i++ {
		_, err := coon.Write([]byte("hello server!"))
		if err != nil {
			fmt.Printf("send data failed, err: %v\n", err)
			break
		}

		result := make([]byte, 1024)
		n, remoteAddr, err := coon.ReadFromUDP(result)

		if err != nil {
			fmt.Printf("receive data failed, err: %v\n", err)
			return
		}

		fmt.Printf("receive data from addr: %v, data: %v\n", remoteAddr, string(result[:n]))
	}
}