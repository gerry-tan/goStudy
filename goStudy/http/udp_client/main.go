package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("connect failed, ", err)
		return
	}

	defer conn.Close()
	//send message
	sendMsg := []byte("hello server!")
	_, err = conn.Write(sendMsg)
	if err != nil {
		fmt.Println("send message failed, ", err)
		return
	}

	//recieve message
	data := make([]byte, 4096)
	len, addr, err := conn.ReadFromUDP(data[:])
	if err != nil {
		fmt.Println("recieve message failed, ", err)
		return
	}
	fmt.Printf("recieve data:%v, addr:%v", string(data[:len]), addr)

}
