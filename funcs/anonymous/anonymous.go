package main

import "fmt"

func main() {
	fmt.Println("This is main function!!")
	x, y := 1, 2
	sum := func(a, b int) int {
		return a + b
	}

	result := sum(x, y)
	fmt.Println("result:", result)
}
