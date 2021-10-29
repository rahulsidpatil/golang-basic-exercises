package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mutex sync.Mutex

func main() {
	done := make(chan struct{})
	for i := 0; i < 3; i++ {
		go routine(done)
	}
	fmt.Println(<-done)
	fmt.Println("I am main!! My id is:", getGID())
	fmt.Println("I am main!! final count is:", count)
}

func routine(done chan<- struct{}) {
	defer func(done chan<- struct{}) {
		done <- struct{}{}
	}(done)

	for i := 0; i < 3; i++ {
		mutex.Lock()
		if count >= 5 {
			mutex.Unlock()
			return
		}
		count++
		fmt.Printf("goroutine:[%v]; count:%v\n", getGID(), count)
		mutex.Unlock()
		time.Sleep(1 * time.Second)
	}
}
