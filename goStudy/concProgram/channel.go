package main

import "fmt"

const num = iota

func Count(ch chan int) {
	//fmt.Println("Counting")
	ch <- num
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for i, ch := range chs {
		number := <-ch
		fmt.Println(i, number)
	}
}
