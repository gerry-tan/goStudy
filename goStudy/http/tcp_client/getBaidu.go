package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, e := net.Dial("tcp", "www.baidu.com:80")
	if e != nil {
		return
	}

	msg := "GET / HTTP/1.1\r\n"
	msg += "HOST: www.baidu.com\r\n"
	msg += "connect: close\r\n"
	msg += "\r\n\r\n"

	_, e = io.WriteString(conn, msg)
	if e != nil {
		return
	}

	var buf [1024]byte
	//for {
	len, _ := conn.Read(buf[:])
	fmt.Println(string(buf[:len]))
	//}

}
