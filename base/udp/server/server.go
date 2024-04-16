package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0,0,0,0),
		Port: 8080,
	})

	if err != nil {
		fmt.Printf("listen failed, err: %v\n", err)
		return
	}

	for {
		var data [1024]byte

		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read failed from addr: %v, err: %v\n", addr, err)
			return
		}

		go func ()  {
			fmt.Printf("addr: %v, data: %v, count: %v\n", addr, string(data[:n]), n)
			_, err := listen.WriteToUDP([]byte("receive success!"), addr)
			if err != nil {
				fmt.Printf("write failed, err: %v \n", err)
			}
		}()
	}
}