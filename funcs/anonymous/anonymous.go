package main

import "fmt"

func main() {
	fmt.Println("This is main function!!")
	sum := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(sum)
}

var sum = func(a, b int) int {
	return a + b
}
