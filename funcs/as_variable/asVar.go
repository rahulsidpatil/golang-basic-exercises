package main

import "fmt"

func main() {
	fmt.Println("This is main function!!")
	fmt.Println(sum(1, 2))
}

var sum = func(a, b int) int {
	return a + b
}
