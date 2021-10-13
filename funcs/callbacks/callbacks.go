package main

import "fmt"

func main() {
	x, y := 1, 2
	subA := func(a, b int) int {
		return b - a
	}
	subB := func(a, b int) int {
		return a - b
	}
	result := doPositiveMath(x, y, subA, subB)
	fmt.Println("result:", result)
}

func doPositiveMath(a, b int, callback1, callback2 func(a, b int) int) int {
	if a > b {
		return callback2(a, b)
	} else {
		return callback1(a, b)
	}
}
