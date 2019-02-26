package main

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"strings"
)

var producer *nsq.Producer

func initProducer(address string) error {
	config := nsq.NewConfig()
	var err error
	producer, err = nsq.NewProducer(address, config)

	if err != nil {
		return err
	}
	return nil
}

func main() {
	var nsqAddress = "127.0.0.1:4150"

	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Printf("init producer failed, err: %v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			continue
		}

		data := strings.TrimSpace(str)

		if data == "stop" {
			break
		}

		err = producer.Publish("test", []byte(data))
		if err != nil {
			fmt.Printf("message send failed, err: %v\n", err)
			continue
		}
		fmt.Println("message send success!")
	}

}
