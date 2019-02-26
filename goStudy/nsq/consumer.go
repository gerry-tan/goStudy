package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Consumer struct {
}

func (*Consumer) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive: ", msg.NSQDAddress, "msg: ", string(msg.Body))
	return nil
}

func initConsumer(address string, topic string, ch string) error {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	var err error

	c, err := nsq.NewConsumer(topic, ch, config)
	if err != nil {
		return err
	}

	consumer := &Consumer{}
	c.AddHandler(consumer)

	err = c.ConnectToNSQLookupd(address)
	return err
}

func consumer(wg *sync.WaitGroup, address string, topic string, ch string) {
	err := initConsumer(address, topic, ch)
	if err != nil {
		fmt.Printf("init consumer failed, err: %v\n", err)
		return
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c

	wg.Done()
}

func main() {
	nsqAddress := "127.0.0.1:4161"

	var wg sync.WaitGroup

	wg.Add(2)

	go consumer(&wg, nsqAddress, "test", "consumer1")
	go consumer(&wg, nsqAddress, "test", "consumer1")

	wg.Wait()

}
