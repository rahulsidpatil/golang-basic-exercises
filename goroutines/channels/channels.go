package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)
	routine(ch)

	for msg := range ch {
		fmt.Printf("%v:Main: Message from routine:%v\n", time.Now(), msg)
	}
}

func routine(ch chan<- string) {
	fmt.Printf("%v:routine: writing to channel.\n", time.Now())
	ch <- "Hello main!!"
	ch <- "how are you??"
	// close(ch)
	// fmt.Printf("%v:routine: rest of the routine.\n", time.Now())
	//....
	//....
}
