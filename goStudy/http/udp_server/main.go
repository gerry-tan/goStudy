package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})

	if err != nil {
		fmt.Println("listen failed, ", err)
		return
	}
	var data [1024]byte
	for {
		len, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read failed, ", err)
			continue
		}
		fmt.Printf("recieve addr:%v, data:%v, length:%v\n", addr, string(data[:len]), len)

		_, err = listen.WriteToUDP([]byte("hello client!"), addr)
		if err != nil {
			fmt.Println("write failed, ", err)
			continue
		}
	}

}
