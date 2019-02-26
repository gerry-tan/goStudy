package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func sendMsg(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)

	for {
		data, e := reader.ReadString('\n')
		if e != nil {
			fmt.Println(e)
			break
		}

		data = strings.TrimSpace(data)

		conn.Write([]byte(data))
	}
}

func main() {
	conn, e := net.Dial("tcp", "localhost:8888")
	if e != nil {
		fmt.Println("dial failed: ", e)
		return
	}

	sendMsg(conn)

}
