package main

import (
	"fmt"
	"net"
)

func process_client(conn net.Conn) {
	for {
		var bytes []byte = make([]byte, 1024)
		n, err := conn.Read(bytes)
		// 从客户端读取数据，阻塞。返回读取的字节数
		if err != nil {
			fmt.Println("Read from client error:", err)
			fmt.Println("Connection with ", conn.RemoteAddr().String(), " down")
			break
		}

		fmt.Println(string(bytes[:n])) // 字节切片转string
	}
}

func main() {
	fmt.Println("Server on..")
	listen, err := net.Listen("tcp", "0.0.0.0:9999")
	// 建立tcp的监听套接字，监听本地9999号端口
	if err != nil {
		fmt.Println("Server listen error..")
		return
	}

	defer listen.Close()

	for {
		fmt.Println("Waiting for client to connect..")
		conn, err := listen.Accept() // 等待客户端连接

		if err != nil {
			fmt.Println("Client connect error..")
			continue
		}

		defer conn.Close()

		fmt.Println("Connection established with ip:", conn.RemoteAddr().String()) // 获取远程地址
		go process_client(conn)

	}

}
