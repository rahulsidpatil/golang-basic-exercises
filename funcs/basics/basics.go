package main

import "fmt"

func main() {
	fmt.Println("This is main function!!")
	fmt.Println(myFunc(1, 2))
}

func myFunc(a, b int) (sum int) {
	return a + b
}
