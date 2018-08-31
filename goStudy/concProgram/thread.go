package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int

func count1(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go count1(lock)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()

		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}
