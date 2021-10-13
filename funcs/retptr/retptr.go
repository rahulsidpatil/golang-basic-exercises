package main

import "fmt"

func main() {
	fmt.Println("result:", *retptr(2))
}

func retptr(x int) *int {
	y := x * x
	return &y
}
