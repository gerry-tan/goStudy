package main

import (
	"fmt"
	"sync"
	"time"
)

func Count(ch chan int) {
	fmt.Println("Counting...")
	ch <- 1 //赋值给变量 ch
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("start goroutine ### ", i)
	time.Sleep(time.Second)
	fmt.Println("end goroutine ### ", i)
	wg.Done()
}

func test1() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i]) //使用go关键字创建一个协程
	}

	//time.Sleep(time.Second)  //等待协程运行结束
	//rang遍历chan队列，优雅的等待协程结束
	for _, ch := range chs {
		<-ch
	}
	fmt.Println("All goroutine finished!")
}

func test2() {
	var wg sync.WaitGroup
	wg.Wait()
	fmt.Println("wait return")

	for i := 0; i < 3; i++ {
		wg.Add(1) //与 wg.done()对应
		go process(i, &wg)
	}
	wg.Wait() //等待结束
	fmt.Println("All goroutine finished!")
}

func main() {
	test2()
	fmt.Println("End.")
}
