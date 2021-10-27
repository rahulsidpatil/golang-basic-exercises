package main

import (
	"fmt"
	"sync"
	"time"
)

var count int

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go routine(&wg)
	}
	wg.Wait()

	fmt.Println("I am main!! My id is:", getGID())
	fmt.Println("I am main!! final count is:", count)
}

func routine(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		count++
		fmt.Println("I am goroutine!! My id is: ", getGID())
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}
