package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go routine(ch)
	go routine1(ch)
	time.Sleep(5 * time.Second)

	fmt.Printf("%v:Main: Message from routine:%v\n", time.Now(), <-ch)
	time.Sleep(5 * time.Second)
}

func routine1(ch chan<- string) {
	ch <- "hey"
	fmt.Printf("%v:Routine1:", time.Now())
	fmt.Printf("%v:Routine1: rest of the routine1.\n", time.Now())
	//....
	//...
}

func routine(ch chan<- string) {
	fmt.Printf("%v:routine: writing to channel.\n", time.Now())
	ch <- "Hello main!!"
	fmt.Printf("%v:routine: rest of the routine.\n", time.Now())
	//....
	//...
}
