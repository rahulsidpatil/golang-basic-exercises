package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go r1(c1)
	go r2(c2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func r1(c1 chan string) {
	time.Sleep(1 * time.Second)
	c1 <- "one"
}

func r2(c2 chan string) {
	time.Sleep(2 * time.Second)
	c2 <- "two"
}
