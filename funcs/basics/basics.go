package main

import "fmt"

func init() {
	fmt.Println("I'm also init func!! I'll go first")
}

func main() {
	fmt.Println("I am main function!!")
	// fmt.Println(myFunc(1, 2))
}

// func myFunc(a, b int) int {
// 	sum := a + b
// 	return sum
// }

func init() {
	fmt.Println("I'm init func!! I'll go first")
}
