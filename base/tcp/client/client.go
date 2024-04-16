package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8081")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err: %v\n", err)
		return
	}

	// 读取命令行输入
	inputReader := bufio.NewReader(os.Stdin)

	for {
		// 一直读到\n
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err: %v\n", err)
			break
		}

		// 读取到Q结束
		trimmedInput := strings.TrimSpace(input) 
		if trimmedInput == "Q" {
			break
		}

		// 回复服务端信息
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Printf("write failed, err: %v\n", err)
			break
		}
	}
}