package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mutex sync.Mutex

func main() {
	countChan := make(chan int)
	for i := 0; i < 3; i++ {
		go increment(countChan)
	}

	for c := range countChan {
		fmt.Printf("main:; current count:%v\n", c)
	}

	fmt.Println("final value of the counter:", count)
}

func increment(countChan chan int) {
	defer close(countChan)
	for {
		mutex.Lock()
		if count >= 5 {
			mutex.Unlock()
			return
		}
		count++
		countChan <- count
		mutex.Unlock()
		time.Sleep(1 * time.Second)
	}
}
