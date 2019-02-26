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
			fmt.Println("read from conn failed, ", err)
			break
		}
		str := string(buf[:n])
		fmt.Println("receive from client, data: ", str)
	}
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen failed: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, ", err)
			continue
		}

		go process(conn)
	}
}
